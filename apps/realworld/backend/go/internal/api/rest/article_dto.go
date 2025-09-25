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

type NewCommentRequest struct {
	Comment *NewComment `json:"comment"`
}

type NewComment struct {
	Body string `json:"body"`
}

type MultiCommentResponse struct {
	Comments []*Comment `json:"comments"`
}

type SingleCommentResponse struct {
	Comment *Comment `json:"comment"`
}

type Comment struct {
	Id        int64     `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Author    *Profile  `json:"author"`
}

func toComment(comment *article.Comment) *Comment {
	return &Comment{
		Id:        comment.Id,
		Body:      comment.Body,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}
