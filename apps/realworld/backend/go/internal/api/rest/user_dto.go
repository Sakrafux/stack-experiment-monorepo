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

func fromNewUser(u *NewUser) *user.User {
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

func toUser(u *user.User) *User {
	return &User{
		Username: u.Username,
		Email:    u.Email,
		Bio:      u.Bio,
		Image:    u.Image,
	}
}

type LoginUserRequest struct {
	User *LoginUser `json:"user"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func fromLoginUser(u *LoginUser) *user.User {
	return &user.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

type UpdateUserRequest struct {
	User *UpdateUser `json:"user"`
}

type UpdateUser struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
	Bio      *string `json:"bio"`
	Image    *string `json:"image"`
}

func fromUpdateUser(u *UpdateUser) *user.UpdateUser {
	return &user.UpdateUser{
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
		Bio:      u.Bio,
		Image:    u.Image,
	}
}
