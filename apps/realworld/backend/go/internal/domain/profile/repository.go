package profile

import "context"

type Repository interface {
	FindProfileById(ctx context.Context, id int64) (*Profile, error)
	FindProfileByIds(ctx context.Context, sourceId, targetId int64) (*Profile, error)
	FollowProfileByIds(ctx context.Context, sourceId, targetId int64) (*Profile, error)
	UnfollowProfileByIds(ctx context.Context, sourceId, targetId int64) (*Profile, error)
}
