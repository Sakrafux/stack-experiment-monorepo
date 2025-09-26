package user

import "context"

type Repository interface {
	Insert(ctx context.Context, user *User) *User
	ExistsByUsernameOrEmail(ctx context.Context, username, email string) bool
	FindByEmail(ctx context.Context, email string) *User
	FindByUsername(ctx context.Context, username string) *User
	FindById(ctx context.Context, id int64) *User
	Update(ctx context.Context, user *User) *User
}
