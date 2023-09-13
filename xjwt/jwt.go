package xjwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(claims jwt.Claims, key string) (string, error) {
	if len(key) < 16 {
		return "", errors.New("key length less 16")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(key))
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

func ParseToken(tokenStr string, key string, claims jwt.Claims) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(*jwt.Token) (any, error) {
		return []byte(key), nil
	})
	if err != nil {
		return token, err
	}
	if !token.Valid {
		return token, errors.New("invalid token")
	}
	return token, nil
}
