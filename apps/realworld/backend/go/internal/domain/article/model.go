package article

import "time"

type Article struct {
	Slug           string
	Title          string
	Description    string
	Body           string
	TagList        []string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Favorited      bool
	FavoritesCount int
	AuthorId       int64
}

type NewArticle struct {
	Slug        *string
	Title       *string
	Description *string
	Body        *string
	TagList     []string
	AuthorId    int64
}

type FilterParams struct {
	Tag       *string
	Author    *string
	Favorited *string
	Offset    int
	Limit     int
	UserId    *int64
}
