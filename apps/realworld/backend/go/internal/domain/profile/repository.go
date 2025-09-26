package profile

import "context"

type Repository interface {
	FindProfileById(ctx context.Context, id int64) *Profile
	FindProfileByIds(ctx context.Context, sourceId, targetId int64) *Profile
	FollowProfileByIds(ctx context.Context, sourceId, targetId int64) *Profile
	UnfollowProfileByIds(ctx context.Context, sourceId, targetId int64)
	FindAllProfilesById(ctx context.Context, ids []int64, userId *int64) []*Profile
}
