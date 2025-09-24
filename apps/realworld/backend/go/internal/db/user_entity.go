package db

import (
	"time"

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

func fromUser(user *user.User) *UserRecord {
	return &UserRecord{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Bio:      user.Bio,
		Image:    user.Image,
	}
}

func toUser(userRecord *UserRecord) *user.User {
	return &user.User{
		Username: userRecord.Username,
		Email:    userRecord.Email,
		Bio:      userRecord.Bio,
		Image:    userRecord.Image,
	}
}
