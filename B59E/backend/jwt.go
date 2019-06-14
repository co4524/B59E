package main

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
	ID       string `json:"id"`
	password string `json:"password"`
	jwt.StandardClaims
}

// Error Message
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "newtrekWang"
)

var jwtSecret = []byte("BCP S3cr3t K3y XD")

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	const duration = 60 * 60 * 12 * time.Hour
	expireTime := nowTime.Add(duration)

	auth := Auth{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, auth)
	token, err := authToken.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Auth, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Auth{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Auth); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
