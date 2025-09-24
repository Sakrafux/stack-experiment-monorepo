package user

import "context"

type Repository interface {
	Insert(ctx context.Context, user *User) (*User, error)
	ExistsByUsernameOrEmail(ctx context.Context, username, email string) (bool, error)
}
