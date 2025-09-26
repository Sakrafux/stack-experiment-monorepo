package user

import (
	"context"
	"fmt"

	"github.com/Sakrafux/stack-experiment-monorepo/pkg/validation"
)

func (s *Service) validateRegisterUser(ctx context.Context, user *User) error {
	exists := s.repo.ExistsByUsernameOrEmail(ctx, user.Username, user.Email)
	if exists {
		return validation.NewValidationError([]string{fmt.Sprintf("The user with username '%s' or email '%s' already exists", user.Username, user.Email)})
	}
	return nil
}

func (s *Service) validateUpdateUser(ctx context.Context, user *UpdateUser) error {
	validations := make([]string, 0)
	if user.Username != nil {
		dbUser := s.repo.FindByUsername(ctx, *user.Username)
		if dbUser != nil && dbUser.Id != user.Id {
			validations = append(validations, fmt.Sprintf("The username '%s' already exists", *user.Username))
		}
	}
	if user.Email != nil {
		dbUser := s.repo.FindByEmail(ctx, *user.Email)
		if dbUser != nil && dbUser.Id != user.Id {
			validations = append(validations, fmt.Sprintf("The email '%s' already exists", *user.Email))
		}
	}

	if len(validations) > 0 {
		return validation.NewValidationError(validations)
	}
	return nil
}
