package rest

import (
	"time"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/domain/article"
)

type MultipleArticlesResponse struct {
	Articles      []*Article `json:"articles"`
	ArticlesCount int        `json:"articlesCount"`
}

type SingleArticleResponse struct {
	Article *Article `json:"article"`
}

type Article struct {
	Slug           string    `json:"slug"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Body           string    `json:"body"`
	TagList        []string  `json:"tagList"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Favorited      bool      `json:"favorited"`
	FavoritesCount int       `json:"favoritesCount"`
	Author         *Profile  `json:"author"`
}

func toArticle(article *article.Article) *Article {
	return &Article{
		Slug:           article.Slug,
		Title:          article.Title,
		Description:    article.Description,
		Body:           article.Body,
		TagList:        article.TagList,
		CreatedAt:      article.CreatedAt,
		UpdatedAt:      article.UpdatedAt,
		Favorited:      article.Favorited,
		FavoritesCount: article.FavoritesCount,
	}
}

type NewArticleRequest struct {
	Article *NewArticle `json:"article"`
}

type NewArticle struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Body        string   `json:"body"`
	TagList     []string `json:"tagList"`
}

func fromNewArticle(newArticle *NewArticle) *article.NewArticle {
	tagList := newArticle.TagList
	if tagList == nil {
		tagList = []string{}
	}
	return &article.NewArticle{
		Title:       &newArticle.Title,
		Description: &newArticle.Description,
		Body:        &newArticle.Body,
		TagList:     tagList,
	}
}

type UpdateArticleRequest struct {
	Article *UpdateArticle `json:"article"`
}

type UpdateArticle struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Body        *string `json:"body"`
}

func fromUpdateArticle(updateArticle *UpdateArticle) *article.NewArticle {
	return &article.NewArticle{
		Title:       updateArticle.Title,
		Description: updateArticle.Description,
		Body:        updateArticle.Body,
	}
}

type TagsResponse struct {
	Tags []string `json:"tags"`
}
