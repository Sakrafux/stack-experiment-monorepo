package article

import "context"

type Repository interface {
	FindAllTags(ctx context.Context) []string
	InsertArticle(ctx context.Context, article *NewArticle) *Article
	FindArticle(ctx context.Context, slug string) *Article
	FindArticleForUser(ctx context.Context, slug string, userId int64) *Article
	UpdateArticle(ctx context.Context, article *NewArticle) *Article
	DeleteArticle(ctx context.Context, slug string)
	FindAllArticlesFiltered(ctx context.Context, filter *FilterParams) []*Article
	FindAllArticlesFeed(ctx context.Context, filter *FilterParams) []*Article
	CreateArticleFavorite(ctx context.Context, slug string, userId int64)
	DeleteArticleFavorite(ctx context.Context, slug string, userId int64)
	FindAllCommentsForArticle(ctx context.Context, slug string) []*Comment
	CreateArticleComment(ctx context.Context, slug string, userId int64, body string) *Comment
	DeleteArticleComment(ctx context.Context, slug string, userId, id int64)
}
