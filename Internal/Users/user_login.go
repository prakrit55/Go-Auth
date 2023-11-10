package user

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	util "github.com/prakrit55/Go-Chat/Util"
)

type LoginReq struct {
	Phone    string `json:"phone" db:"phone"`
	Password string `json:"password" db:"password"`
}

type LoginRes struct {
	accessToken string
	ID          string `json:"id" db:"id"`
	Username    string `json:"username" db:"username"`
}

type Service interface {
	CreateUser(c context.Context, req *UserReq) (*UserRes, error)
	Login(c context.Context, req *LoginReq) (*LoginRes, error)
}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// User uses email and password to login
func (s *service) Login(c context.Context, req *LoginReq) (*LoginRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repo.GetUserByPhone(ctx, req.Phone)
	if err != nil {
		return &LoginRes{}, err
	}

	// Decrypt the password of user and checks the user is present or not
	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginRes{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginRes{}, err
	}

	return &LoginRes{accessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID))}, nil
}
