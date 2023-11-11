package user

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	jwtclaim "github.com/prakrit55/Go-Auth/JWT"
	util "github.com/prakrit55/Go-Auth/Util"
)

type DataReq struct {
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
}

type Service interface {
	CreateUser(c context.Context, req *UserReq) error
	Login(c context.Context, req *DataReq) (string, error)
	DeleteUser(c context.Context, req *DataReq) error
}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// User uses email and password to login
func (s *service) Login(c context.Context, req *DataReq) (string, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repo.GetUserByPhone(ctx, req.Phone)
	if err != nil {
		return "", err
	}

	// Decrypt the password of user and checks the user is present or not
	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return "", err
	}

	ss, err := jwtclaim.GenerateJWT(u.ID, u.Username)
	if err != nil {
		return "", err
	}

	return ss, nil
}
