package user

import (
	"context"
	"fmt"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/config"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/errors"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/security"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/util"
)

type Service struct {
	config *config.Config
	repo   Repository
}

func NewService(config *config.Config, repo Repository) *Service {
	return &Service{config, repo}
}

func (s *Service) RegisterUser(ctx context.Context, user *User) (*User, error) {
	err := s.validateRegisterUser(ctx, user)
	if err != nil {
		return nil, err
	}

	password, err := security.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = password

	u := s.repo.Insert(ctx, user)
	if u == nil {
		return nil, errors.NewConflictError("user not created")
	}
	return u, nil
}

func (s *Service) LoginUser(ctx context.Context, user *User) (*User, error) {
	u := s.repo.FindByEmail(ctx, user.Email)
	if u == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("user with email '%s' not found", user.Email))
	}

	if security.CheckPassword(u.Password, user.Password) {
		return u, nil
	}

	return nil, errors.NewUnauthorizedError("invalid user")
}

func (s *Service) FindUserById(ctx context.Context, id int64) (*User, error) {
	u := s.repo.FindById(ctx, id)
	if u == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("user with id '%d' not found", id))
	}
	return u, nil
}

func (s *Service) UpdateUser(ctx context.Context, updateUser *UpdateUser) (*User, error) {
	err := s.validateUpdateUser(ctx, updateUser)
	if err != nil {
		return nil, err
	}

	user, err := s.FindUserById(ctx, updateUser.Id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("user with id '%d' not found", updateUser.Id))
	}

	user.Username = util.DerefOrDefault(updateUser.Username, user.Username)
	user.Email = util.DerefOrDefault(updateUser.Email, user.Email)
	user.Bio = util.DerefOrDefault(updateUser.Bio, user.Bio)
	user.Image = util.DerefOrDefault(updateUser.Image, user.Image)

	if updateUser.Password != nil {
		password, err := security.HashPassword(*updateUser.Password)
		if err != nil {
			return nil, err
		}
		user.Password = password
	}

	u := s.repo.Update(ctx, user)
	if u == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("user with id '%d' not found", updateUser.Id))
	}
	return u, nil
}
