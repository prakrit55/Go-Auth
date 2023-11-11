package user

import (
	"context"
	"time"

	util "github.com/prakrit55/Go-Auth/Util"
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

func (s *service) CreateUser(c context.Context, req *UserReq) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// Encrypt the password in Util
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashedPassword,
	}

	// Stores the user data in db with the hashed password
	err = s.Repo.CreateUser(ctx, u)
	if err != nil {
		return err
	}
	
	return nil
}

func (s *service) DeleteUser(c context.Context, req *DataReq) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	err := s.Repo.DeleteUserByPhone(ctx, req.Phone)
	if err != nil {
		return err
	}
	return nil
}
