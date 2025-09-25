package db

import (
	"time"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/domain/article"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/util"
)

type ArticleRecord struct {
	Id          int64     `db:"id"`
	Slug        string    `db:"slug"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Body        string    `db:"body"`
	AuthorId    int64     `db:"fk_author"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Version     int       `db:"version"`
}

type TagRecord struct {
	Id        int64     `db:"id"`
	Tag       string    `db:"tag"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Version   int       `db:"version"`
}

type TagJoinRecord struct {
	ArticleId int64 `db:"article_id"`
	TagId     int64 `db:"tag_id"`
}

func fromNewArticle(a *article.NewArticle) *ArticleRecord {
	return &ArticleRecord{
		Slug:        util.DerefOrDefault(a.Slug, ""),
		Title:       util.DerefOrDefault(a.Title, ""),
		Description: util.DerefOrDefault(a.Description, ""),
		Body:        util.DerefOrDefault(a.Body, ""),
		AuthorId:    a.AuthorId,
	}
}

func toArticle(a *ArticleRecord) *article.Article {
	return &article.Article{
		Slug:        a.Slug,
		Title:       a.Title,
		Description: a.Description,
		Body:        a.Body,
		AuthorId:    a.AuthorId,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}
