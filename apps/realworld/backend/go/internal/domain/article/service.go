package article

import (
	"context"
	"strings"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/config"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/errors"
)

type Service struct {
	config *config.Config
	repo   Repository
}

func NewService(config *config.Config, repo Repository) *Service {
	return &Service{config, repo}
}

func (s *Service) CreateArticle(ctx context.Context, article *NewArticle) (*Article, error) {
	slug := strings.ReplaceAll(strings.ToLower(*article.Title), " ", "-")
	article.Slug = &slug
	err := s.validateCreateArticle(ctx, article)
	if err != nil {
		return nil, err
	}

	a := s.repo.InsertArticle(ctx, article)
	if a == nil {
		return nil, errors.NewConflictError("article not updated")
	}
	return a, nil
}

func (s *Service) GetArticle(ctx context.Context, slug string, userId int64) (*Article, error) {
	if userId < 0 {
		a := s.repo.FindArticle(ctx, slug)
		if a == nil {
			return nil, errors.NewNotFoundError("article not found")
		}
		return a, nil
	}
	a := s.repo.FindArticleForUser(ctx, slug, userId)
	if a == nil {
		return nil, errors.NewNotFoundError("article not found")
	}
	return a, nil
}

func (s *Service) UpdateArticle(ctx context.Context, article *NewArticle) (*Article, error) {
	err := s.validateUpdateArticle(ctx, article)
	if err != nil {
		return nil, err
	}

	a := s.repo.FindArticle(ctx, *article.Slug)
	if a == nil {
		return nil, errors.NewNotFoundError("article not found")
	}

	if article.Title == nil {
		article.Title = &a.Title
	}
	if article.Body == nil {
		article.Body = &a.Body
	}
	if article.Description == nil {
		article.Description = &a.Description
	}

	a = s.repo.UpdateArticle(ctx, article)
	if a == nil {
		return nil, errors.NewConflictError("article not updated")
	}
	return a, nil
}

func (s *Service) DeleteArticle(ctx context.Context, slug string) error {
	err := s.validateDeleteArticle(ctx, slug)
	if err != nil {
		return err
	}

	s.repo.DeleteArticle(ctx, slug)
	return nil
}

func (s *Service) GetTags(ctx context.Context) []string {
	return s.repo.FindAllTags(ctx)
}

func (s *Service) GetArticles(ctx context.Context, filter *FilterParams) []*Article {
	return s.repo.FindAllArticlesFiltered(ctx, filter)
}

func (s *Service) GetArticlesFeed(ctx context.Context, filter *FilterParams) []*Article {
	return s.repo.FindAllArticlesFeed(ctx, filter)
}

func (s *Service) CreateArticleFavorite(ctx context.Context, slug string, userId int64) {
	s.repo.CreateArticleFavorite(ctx, slug, userId)
}

func (s *Service) DeleteArticleFavorite(ctx context.Context, slug string, userId int64) {
	s.repo.DeleteArticleFavorite(ctx, slug, userId)
}

func (s *Service) GetArticleComments(ctx context.Context, slug string) ([]*Comment, error) {
	err := s.validateArticleExists(ctx, slug)
	if err != nil {
		return nil, err
	}

	return s.repo.FindAllCommentsForArticle(ctx, slug), nil
}

func (s *Service) CreateArticleComment(ctx context.Context, slug string, userId int64, body string) (*Comment, error) {
	err := s.validateArticleExists(ctx, slug)
	if err != nil {
		return nil, err
	}

	c := s.repo.CreateArticleComment(ctx, slug, userId, body)
	if c == nil {
		return nil, errors.NewConflictError("comment not created")
	}
	return c, nil
}

func (s *Service) DeleteArticleComment(ctx context.Context, slug string, userId, id int64) error {
	err := s.validateArticleExists(ctx, slug)
	if err != nil {
		return err
	}

	s.repo.DeleteArticleComment(ctx, slug, userId, id)
	return nil
}
