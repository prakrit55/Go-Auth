package user

import "context"

type User struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
}

type UserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Repo interface {
	CreateUser(ctx context.Context, user *User) (error)
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
	DeleteUserByPhone(ctx context.Context, phone string) (error)
}
