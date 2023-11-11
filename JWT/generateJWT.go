package jwtclaim

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	secretKey = "secret"
)

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(ID int64, Email string) (ss string, err error) {
	jtClaim := &MyJWTClaims{
		ID:       strconv.Itoa(int(ID)),
		Username: Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jtClaim)
	ss, err = newToken.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}
	return ss, nil
}
