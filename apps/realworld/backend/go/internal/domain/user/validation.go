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
