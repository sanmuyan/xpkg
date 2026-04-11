package xjwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

// ValidJwtKey 校验 JWT 密钥长度是否符合要求
func ValidJWTKey(key []byte) error {
	switch len(key) {
	case 32:
		return nil
	default:
		return errors.New("key length invalid")
	}
}

// CreateToken 创建 JWT
func CreateToken(claims jwt.Claims, key []byte) (tokenStr string, err error) {
	err = ValidJWTKey(key)
	if err != nil {
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(key)
	if err != nil {
		return
	}
	return
}

// ParseToken 解析 JWT
func ParseToken(tokenStr string, key []byte, claims jwt.Claims) (token *jwt.Token, err error) {
	err = ValidJWTKey(key)
	if err != nil {
		return
	}
	token, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected method")
		}
		return key, nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		err = errors.New("token is invalid")
		return
	}
	return
}

// CreateTokenRSA 创建 JWT，需要提供 PEM 私钥
func CreateTokenRSA(claims jwt.Claims, privateKey []byte) (tokenStr string, err error) {
	pk, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenStr, err = token.SignedString(pk)
	if err != nil {
		return
	}
	return
}

// ParseTokenRSA 解析 JWT，需要提供 PEM 公钥
func ParseTokenRSA(tokenStr string, publicKey []byte, claims jwt.Claims) (token *jwt.Token, err error) {
	pk, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return
	}
	token, err = jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected method")
		}
		return pk, nil
	})
	if err != nil {
		return
	}
	if !token.Valid {
		err = errors.New("invalid token")
		return
	}
	return
}
