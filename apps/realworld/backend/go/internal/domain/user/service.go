package user

import (
	"context"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/config"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/errors"
	"github.com/Sakrafux/stack-experiment-monorepo/pkg/security"
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

	return s.repo.Insert(ctx, user)
}

func (s *Service) LoginUser(ctx context.Context, user *User) (*User, error) {
	u, err := s.repo.FindByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if security.CheckPassword(u.Password, user.Password) {
		return u, nil
	}

	return nil, errors.NewUnauthorizedError("invalid user")
}
