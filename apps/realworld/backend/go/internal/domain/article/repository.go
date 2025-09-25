package article

import "context"

type Repository interface {
	FindAllTags(ctx context.Context) ([]string, error)
	InsertArticle(ctx context.Context, article *NewArticle) (*Article, error)
	FindArticle(ctx context.Context, slug string) (*Article, error)
	FindArticleForUser(ctx context.Context, slug string, userId int64) (*Article, error)
	UpdateArticle(ctx context.Context, article *NewArticle) (*Article, error)
	DeleteArticle(ctx context.Context, slug string) error
	FindAllArticlesFiltered(ctx context.Context, filter *FilterParams) ([]*Article, error)
	FindAllArticlesFeed(ctx context.Context, filter *FilterParams) ([]*Article, error)
	CreateArticleFavorite(ctx context.Context, slug string, userId int64) error
	DeleteArticleFavorite(ctx context.Context, slug string, userId int64) error
	FindAllCommentsForArticle(ctx context.Context, slug string) ([]*Comment, error)
	CreateArticleComment(ctx context.Context, slug string, userId int64, body string) (*Comment, error)
	DeleteArticleComment(ctx context.Context, slug string, userId, id int64) error
}
