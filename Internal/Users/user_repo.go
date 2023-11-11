package user

import (
	"context"
	"database/sql"
	"log"
)

type DBfn interface {
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

type repo struct {
	db DBfn
}

func NewRepo(db DBfn) Repo {
	return &repo{db: db}
}

// CreateUser function stores user's data in database and returns the user alongwith the id
func (r *repo) CreateUser(ctx context.Context, user *User) (error) {
	var lastInsertId int
	query := "INSERT INTO users(username, password, email, phone) VALUES ($1, $2, $3, $4) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email, user.Phone).Scan(&lastInsertId)
	if err != nil {
		return err
	}
	user.ID = int64(lastInsertId)
	log.Print("user created", user.ID)
	return nil
}

func (r *repo) GetUserByPhone(ctx context.Context, phone string) (*User, error) {
	u := User{}
	query := "SELECT id, email, username, phone, password FROM users WHERE phone = $1"
	err := r.db.QueryRowContext(ctx, query, phone).Scan(&u.ID, &u.Email, &u.Username, &u.Phone, &u.Password)
	if err != nil {
		return &User{}, err
	}
	log.Print("got user", u.ID)
	return &u, nil
}

func (r *repo) DeleteUserByPhone(ctx context.Context, phone string) error { 
	query := "DELETE FROM users WHERE phone = $1"
	_, err := r.db.ExecContext(ctx, query, phone)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Print("Deleted user")
	return nil
}
