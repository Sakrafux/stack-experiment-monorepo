package db

import (
	"time"

	"github.com/Sakrafux/stack-experiment-monorepo/internal/domain/profile"
	"github.com/Sakrafux/stack-experiment-monorepo/internal/domain/user"
)

type UserRecord struct {
	Id        int64     `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Bio       string    `db:"bio"`
	Image     string    `db:"image"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Version   int       `db:"version"`
}

type FollowRecord struct {
	SourceId int64 `db:"following_user_id"`
	TargetId int64 `db:"followed_user_id"`
}

func fromUser(user *user.User) *UserRecord {
	return &UserRecord{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Bio:      user.Bio,
		Image:    user.Image,
	}
}

func toUser(userRecord *UserRecord) *user.User {
	return &user.User{
		Id:       userRecord.Id,
		Username: userRecord.Username,
		Email:    userRecord.Email,
		Password: userRecord.Password,
		Bio:      userRecord.Bio,
		Image:    userRecord.Image,
	}
}

func toProfile(record *UserRecord) *profile.Profile {
	return &profile.Profile{
		Username: record.Username,
		Bio:      record.Bio,
		Image:    record.Image,
	}
}
