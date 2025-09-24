package user

import (
	"context"

	"github.com/Sakrafux/stack-experiment-monorepo/pkg/security"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo}
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
