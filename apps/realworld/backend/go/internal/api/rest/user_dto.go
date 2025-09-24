package rest

import "github.com/Sakrafux/stack-experiment-monorepo/internal/domain/user"

type NewUserRequest struct {
	User *NewUser `json:"user"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func toUser(u *NewUser) *user.User {
	return &user.User{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}

type UserResponse struct {
	User *User `json:"user"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Token    string `json:"token"`
}

func fromUser(u *user.User) *User {
	return &User{
		Username: u.Username,
		Email:    u.Email,
		Bio:      u.Bio,
		Image:    u.Image,
	}
}
