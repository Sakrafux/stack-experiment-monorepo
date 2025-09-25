package user

import "context"

type Repository interface {
	Insert(ctx context.Context, user *User) (*User, error)
	ExistsByUsernameOrEmail(ctx context.Context, username, email string) (bool, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindById(ctx context.Context, id int64) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
}
