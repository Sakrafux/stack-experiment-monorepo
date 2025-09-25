package user

import (
	"context"
	"fmt"

	"github.com/Sakrafux/stack-experiment-monorepo/pkg/validation"
)

func (s *Service) validateRegisterUser(ctx context.Context, user *User) error {
	exists, err := s.repo.ExistsByUsernameOrEmail(ctx, user.Username, user.Email)
	if err != nil {
		return err
	}
	if exists {
		return validation.NewValidationError([]string{fmt.Sprintf("The user with username '%s' or email '%s' already exists", user.Username, user.Email)})
	}
	return nil
}

func (s *Service) validateUpdateUser(ctx context.Context, user *UpdateUser) error {
	validations := make([]string, 0)
	if user.Username != nil {
		dbUser, err := s.repo.FindByUsername(ctx, *user.Username)
		if err != nil {
			return err
		}
		if dbUser != nil && dbUser.Id != user.Id {
			validations = append(validations, fmt.Sprintf("The username '%s' already exists", *user.Username))
		}
	}
	if user.Email != nil {
		dbUser, err := s.repo.FindByEmail(ctx, *user.Email)
		if err != nil {
			return err
		}
		if dbUser != nil && dbUser.Id != user.Id {
			validations = append(validations, fmt.Sprintf("The email '%s' already exists", *user.Email))
		}
	}

	if len(validations) > 0 {
		return validation.NewValidationError(validations)
	}
	return nil
}
