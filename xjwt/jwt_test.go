package xjwt

import (
	"github.com/golang-jwt/jwt/v4"
	"testing"
)

func TestToken(t *testing.T) {
	token, err := CreateToken(jwt.RegisteredClaims{
		ID: "123456",
	}, "xxxxxxxxxxxxxxxx")
	if err != nil {
		t.Error(err)
	}
	sc := &jwt.RegisteredClaims{}
	_, err = ParseToken(token, "xxxxxxxxxxxxxxxx", sc)
	if err != nil {
		t.Error(err)
	}
	if string(sc.ID) != "123456" {
		t.Error(sc.ID)
	}
	t.Log(sc)
	_, err = ParseToken(token, "xxxxxxxxxxxxxxxa", sc)
	if err == nil {
		t.Error("invalid key but no return error")
	}
}
