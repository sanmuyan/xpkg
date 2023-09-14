package xjwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

// CreateToken 创建 JWT，key 必须大于等于32
func CreateToken(claims jwt.Claims, key string) (string, error) {
	if len(key) < 32 {
		return "", errors.New("key length invalid")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(key))
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

// ParseToken 解析 JWT
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
