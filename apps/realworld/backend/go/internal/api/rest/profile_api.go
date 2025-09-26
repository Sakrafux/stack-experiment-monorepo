package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/db"
	"github.com/Sakrafux/stack-experiment-monorepo/internal/domain/profile"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/errors"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/middleware"
	"github.com/go-chi/chi/v5"
)

type ProfileApi struct {
	api     *Api
	repo    *db.UserRepository
	service *profile.Service
}

func NewProfileApi(api *Api) *ProfileApi {
	repo := db.NewUserRepository(api.db)
	service := profile.NewService(api.config, repo)
	return &ProfileApi{api, repo, service}
}

func (api *ProfileApi) CreateProfilesRouter() http.Handler {
	r := chi.NewRouter()

	r.Route("/{username}", func(r chi.Router) {
		r.Use(api.UsernameCtx)
		r.Get("/", api.GetProfileByUsername)
		r.With(middleware.Authorization()).Post("/follow", api.FollowUserByUsername)
		r.With(middleware.Authorization()).Delete("/follow", api.UnfollowUserByUsername)
	})

	return r
}

func (api *ProfileApi) UsernameCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		username := chi.URLParam(r, "username")

		user, err := api.repo.FindByUsername(ctx, username)
		if err != nil {
			errors.HandleHttpError(w, r, err)
			return
		}
		if user == nil {
			errors.HandleHttpError(w, r, errors.NewNotFoundError(fmt.Sprintf("user '%s' not found", username)))
		}
		ctx = context.WithValue(r.Context(), "targetUserId", user.Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *ProfileApi) GetProfileByUsername(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sourceUserId, ok := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)
	if !ok {
		sourceUserId = -1
	}
	targetUserId := ctx.Value("targetUserId").(int64)

	p := api.service.GetProfile(ctx, sourceUserId, targetUserId)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(ProfileResponse{toProfile(p)})
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
}

func (api *ProfileApi) FollowUserByUsername(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sourceUserId := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)
	targetUserId := ctx.Value("targetUserId").(int64)

	p := api.service.FollowUser(ctx, sourceUserId, targetUserId)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(ProfileResponse{toProfile(p)})
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
}

func (api *ProfileApi) UnfollowUserByUsername(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	sourceUserId := ctx.Value(middleware.AUTH_CONTEXT_ID).(int64)
	targetUserId := ctx.Value("targetUserId").(int64)

	p := api.service.UnfollowUser(ctx, sourceUserId, targetUserId)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(ProfileResponse{toProfile(p)})
	if err != nil {
		errors.HandleHttpError(w, r, err)
		return
	}
}
