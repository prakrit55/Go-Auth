package user

import (
	"context"
	"strconv"
	"time"

	util "github.com/prakrit55/Go-Chat/Util"
)

const (
	secretKey = "secret"
)

type service struct {
	Repo
	timeout time.Duration
}

func NewService(repository Repo) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *UserReq) (*UserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// Encrypt the password in Util
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashedPassword,
	}

	// Stores the user data in db with the hashed password
	r, err := s.Repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &UserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
		Phone:    r.Phone,
	}
	return res, nil
}
