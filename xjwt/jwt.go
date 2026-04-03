package xjwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

// CreateToken 创建 JWT，密钥长度不能小于32
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
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected method")
		}
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

// CreateTokenRSA 创建 JWT，需要提供私钥 PEM 格式字符串
func CreateTokenRSA(claims jwt.Claims, privateKeyStr string) (string, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyStr))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenStr, err := token.SignedString(privateKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

// ParseTokenRSA 解析 JWT，需要提供公钥 PEM 格式字符串
func ParseTokenRSA(tokenStr string, publicKeyStr string, claims jwt.Claims) (*jwt.Token, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyStr))
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected method")
		}
		return publicKey, nil
	})
	if err != nil {
		return token, err
	}
	if !token.Valid {
		return token, errors.New("invalid token")
	}
	return token, nil
}
