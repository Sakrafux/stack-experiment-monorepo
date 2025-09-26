package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/db"
	"github.com/Sakrafux/stack-experiment-monorepo/internal/domain/article"
	"github.com/Sakrafux/stack-experiment-monorepo/internal/domain/profile"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/errors"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/samber/lo"
)

type ArticleApi struct {
	api            *Api
	repo           *db.ArticleRepository
	service        *article.Service
	userRepo       *db.UserRepository
	profileService *profile.Service
}

func NewArticleApi(api *Api) *ArticleApi {
	repo := db.NewArticleRepository(api.db)
	service := article.NewService(api.config, repo)
	userRepo := db.NewUserRepository(api.db)
	profileService := profile.NewService(api.config, userRepo)
	return &ArticleApi{api, repo, service, userRepo, profileService}
}

func (api *ArticleApi) CreateArticlesRouter() http.Handler {
	r := chi.NewRouter()

	r.With(middleware.Authorization()).Get("/feed", api.GetArticlesFeed)
	r.Get("/", api.GetArticles)
	r.Post("/", api.CreateArticle)

	r.Route("/{slug}", func(r chi.Router) {
		r.Use(api.SlugCtx)
		r.Get("/", api.GetArticle)
		r.With(middleware.Authorization()).Put("/", api.UpdateArticle)
		r.With(middleware.Authorization()).Delete("/", api.DeleteArticle)

		r.Route("/favorite", func(r chi.Router) {
			r.Use(middleware.Authorization())
			r.Post("/", api.CreateArticleFavorite)
			r.Delete("/", api.DeleteArticleFavorite)
		})

		r.Route("/comments", func(r chi.Router) {
			r.Get("/", api.GetArticleComments)
			r.With(middleware.Authorization()).Post("/", api.CreateArticleComment)
			r.With(middleware.Authorization()).Delete("/{id}", api.DeleteArticleComment)
		})
	})

	return r
}

func (api *ArticleApi) CreateTagsRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", api.GetTags)

	return r
}

func (api *ArticleApi) SlugCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		slug := chi.URLParam(r, "slug")

		if slug == "" {
			errors.HandleHttpError(w, r, errors.NewBadRequestError("slug required"))
			return
		}

		ctx = context.WithValue(r.Context(), "slug", slug)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *ArticleApi) GetArticlesFeed(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := article.FilterParams{}
	filter.Offset = 0
	filter.Limit = 20
	if r.URL.Query().Get("limit") != "" {
		res, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			errors.HandleHttpError(w, r, err)
			return
		}
		filter.Limit = res
	}
	if r.URL.Query().Get("offset") != "" {
		res, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			errors.HandleHttpError(w, r, err)
			return
		}
		filter.Offset = res
	}
	userId := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)
	filter.UserId = &userId

	feed, err := api.service.GetArticlesFeed(ctx, &filter)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	authorIds := lo.Map(feed, func(item *article.Article, index int) int64 {
		return item.AuthorId
	})
	profiles := api.userRepo.FindAllProfilesById(ctx, authorIds, filter.UserId)
	authors := lo.SliceToMap(profiles, func(item *profile.Profile) (int64, *profile.Profile) {
		return item.Id, item
	})

	dtos := lo.Map(feed, func(item *article.Article, index int) *Article {
		dto := toArticle(item)
		dto.Author = toProfile(authors[item.AuthorId])
		return dto
	})

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(MultipleArticlesResponse{dtos, len(dtos)})
	if err != nil {
		errors.HandleHttpError(w, r, err)
	}
}

func (api *ArticleApi) GetArticles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	filter := article.FilterParams{}
	filter.Offset = 0
	filter.Limit = 20
	if r.URL.Query().Get("limit") != "" {
		res, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			errors.HandleHttpError(w, r, err)
			return
		}
		filter.Limit = res
	}
	if r.URL.Query().Get("offset") != "" {
		res, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			errors.HandleHttpError(w, r, err)
			return
		}
		filter.Offset = res
	}
	if tag := r.URL.Query().Get("tag"); tag != "" {
		filter.Tag = &tag
	}
	if author := r.URL.Query().Get("author"); author != "" {
		filter.Author = &author
	}
	if favorited := r.URL.Query().Get("favorited"); favorited != "" {
		filter.Favorited = &favorited
	}
	userId, ok := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)
	if ok {
		filter.UserId = &userId
	} else {
		userId = -1
	}

	feed, err := api.service.GetArticles(ctx, &filter)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	authorIds := lo.Map(feed, func(item *article.Article, index int) int64 {
		return item.AuthorId
	})
	profiles := api.userRepo.FindAllProfilesById(ctx, authorIds, filter.UserId)
	authors := lo.SliceToMap(profiles, func(item *profile.Profile) (int64, *profile.Profile) {
		return item.Id, item
	})

	dtos := lo.Map(feed, func(item *article.Article, index int) *Article {
		dto := toArticle(item)
		dto.Author = toProfile(authors[item.AuthorId])
		return dto
	})

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(MultipleArticlesResponse{dtos, len(dtos)})
	if err != nil {
		errors.HandleHttpError(w, r, err)
	}
}

func (api *ArticleApi) CreateArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req NewArticleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errors.HandleHttpError(w, r, errors.NewBadRequestError(err.Error()))
		return
	}
	sourceUserId := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)

	input := fromNewArticle(req.Article)
	input.AuthorId = sourceUserId
	a, err := api.service.CreateArticle(ctx, input)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	author := api.profileService.GetProfile(ctx, sourceUserId, a.AuthorId)

	dto := toArticle(a)
	dto.Author = toProfile(author)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(SingleArticleResponse{dto})
	if err != nil {
		errors.HandleHttpError(w, r, err)
	}
}

func (api *ArticleApi) GetArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slug := ctx.Value("slug").(string)
	sourceUserId, ok := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)
	if !ok {
		sourceUserId = -1
	}

	a, err := api.service.GetArticle(ctx, slug, sourceUserId)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	author := api.profileService.GetProfile(ctx, sourceUserId, a.AuthorId)

	dto := toArticle(a)
	dto.Author = toProfile(author)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(SingleArticleResponse{dto})
	if err != nil {
		errors.HandleHttpError(w, r, err)
	}
}

func (api *ArticleApi) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slug := ctx.Value("slug").(string)
	var req UpdateArticleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errors.HandleHttpError(w, r, errors.NewBadRequestError(err.Error()))
		return
	}
	sourceUserId := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)

	input := fromUpdateArticle(req.Article)
	input.Slug = &slug
	input.AuthorId = sourceUserId

	a, err := api.service.UpdateArticle(ctx, input)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	author := api.profileService.GetProfile(ctx, sourceUserId, a.AuthorId)

	dto := toArticle(a)
	dto.Author = toProfile(author)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(SingleArticleResponse{dto})
	if err != nil {
		errors.HandleHttpError(w, r, err)
	}
}

func (api *ArticleApi) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slug := ctx.Value("slug").(string)

	err := api.service.DeleteArticle(ctx, slug)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (api *ArticleApi) GetTags(w http.ResponseWriter, r *http.Request) {
	tags, err := api.service.GetTags(r.Context())
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(TagsResponse{tags})
	if err != nil {
		errors.HandleHttpError(w, r, err)
	}
}

func (api *ArticleApi) CreateArticleFavorite(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)
	slug := ctx.Value("slug").(string)

	err := api.service.CreateArticleFavorite(ctx, slug, userId)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	api.GetArticle(w, r)
}

func (api *ArticleApi) DeleteArticleFavorite(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)
	slug := ctx.Value("slug").(string)

	err := api.service.DeleteArticleFavorite(ctx, slug, userId)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	api.GetArticle(w, r)
}

func (api *ArticleApi) GetArticleComments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slug := ctx.Value("slug").(string)
	userId, ok := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)
	var userIdRef *int64
	if ok {
		userIdRef = &userId
	}

	comments, err := api.service.GetArticleComments(ctx, slug)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	authorIds := lo.Map(comments, func(item *article.Comment, index int) int64 {
		return item.AuthorId
	})
	profiles := api.userRepo.FindAllProfilesById(ctx, authorIds, userIdRef)
	authors := lo.SliceToMap(profiles, func(item *profile.Profile) (int64, *profile.Profile) {
		return item.Id, item
	})

	dtos := lo.Map(comments, func(item *article.Comment, index int) *Comment {
		dto := toComment(item)
		dto.Author = toProfile(authors[item.AuthorId])
		return dto
	})

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(MultiCommentResponse{dtos})
	if err != nil {
		errors.HandleHttpError(w, r, err)
	}
}

func (api *ArticleApi) CreateArticleComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slug := ctx.Value("slug").(string)
	userId := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)

	var req NewCommentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errors.HandleHttpError(w, r, errors.NewBadRequestError(err.Error()))
		return
	}

	comment, err := api.service.CreateArticleComment(ctx, slug, userId, req.Comment.Body)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	dto := toComment(comment)

	author := api.profileService.GetProfile(ctx, userId, comment.AuthorId)

	dto.Author = toProfile(author)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(SingleCommentResponse{dto})
	if err != nil {
		errors.HandleHttpError(w, r, err)
	}
}

func (api *ArticleApi) DeleteArticleComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slug := ctx.Value("slug").(string)
	userId := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}

	err = api.service.DeleteArticleComment(ctx, slug, userId, id)
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
