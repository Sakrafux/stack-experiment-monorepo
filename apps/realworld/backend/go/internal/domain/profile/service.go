package profile

import (
	"context"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/config"
)

type Service struct {
	config *config.Config
	repo   Repository
}

func NewService(config *config.Config, repo Repository) *Service {
	return &Service{config, repo}
}

func (s *Service) GetProfile(ctx context.Context, sourceId, targetId int64) *Profile {
	if sourceId < 0 {
		return s.repo.FindProfileById(ctx, targetId)
	}
	return s.repo.FindProfileByIds(ctx, sourceId, targetId)
}

func (s *Service) FollowUser(ctx context.Context, sourceId, targetId int64) *Profile {
	return s.repo.FollowProfileByIds(ctx, sourceId, targetId)
}

func (s *Service) UnfollowUser(ctx context.Context, sourceId, targetId int64) *Profile {
	s.repo.UnfollowProfileByIds(ctx, sourceId, targetId)
	return s.repo.FindProfileByIds(ctx, sourceId, targetId)
}
