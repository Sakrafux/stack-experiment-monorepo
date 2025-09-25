package article

import (
	"context"

	"github.com/Sakrafux/stack-experiment-monorepo/pkg/util"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/validation"
)

func (s *Service) validateCreateArticle(ctx context.Context, article *NewArticle) error {
	validations := make([]string, 0)

	if article.Title == nil || article.Body == nil || article.Description == nil {
		validations = append(validations, "not all fields are present")
	}

	a, err := s.repo.FindArticle(ctx, util.DerefOrDefault(article.Slug, ""))
	if err != nil {
		return err
	}
	if a != nil {
		validations = append(validations, "slug is already present")
	}

	if len(validations) > 0 {
		return validation.NewValidationError(validations)
	}

	return nil
}

func (s *Service) validateUpdateArticle(ctx context.Context, article *NewArticle) error {
	validations := make([]string, 0)

	if article.Title == nil && article.Body == nil && article.Description == nil {
		validations = append(validations, "no updated field is present")
	}

	a, err := s.repo.FindArticle(ctx, util.DerefOrDefault(article.Slug, ""))
	if err != nil {
		return err
	}
	if a == nil {
		validations = append(validations, "slug does not exist")
	}

	if len(validations) > 0 {
		return validation.NewValidationError(validations)
	}

	return nil
}

func (s *Service) validateDeleteArticle(ctx context.Context, slug string) error {
	validations := make([]string, 0)

	a, err := s.repo.FindArticle(ctx, slug)
	if err != nil {
		return err
	}
	if a == nil {
		validations = append(validations, "slug does not exist")
	}

	if len(validations) > 0 {
		return validation.NewValidationError(validations)
	}

	return nil
}

func (s *Service) validateArticleExists(ctx context.Context, slug string) error {
	validations := make([]string, 0)

	a, err := s.repo.FindArticle(ctx, slug)
	if err != nil {
		return err
	}
	if a == nil {
		validations = append(validations, "slug does not exist")
	}

	if len(validations) > 0 {
		return validation.NewValidationError(validations)
	}

	return nil
}
