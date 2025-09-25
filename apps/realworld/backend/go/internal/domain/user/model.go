package user

type User struct {
	Id       int64
	Username string
	Email    string
	Password string
	Bio      string
	Image    string
}

type UpdateUser struct {
	Id       int64
	Username *string
	Email    *string
	Password *string
	Bio      *string
	Image    *string
}
