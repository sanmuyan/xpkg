package xjwt

import (
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestToken(t *testing.T) {
	key := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	token, err := CreateToken(jwt.RegisteredClaims{
		ID: "123456",
	}, key)
	if err != nil {
		t.Error(err)
	}
	sc := &jwt.RegisteredClaims{}
	_, err = ParseToken(token, key, sc)
	if err != nil {
		t.Error(err)
	}
	if string(sc.ID) != "123456" {
		t.Error(sc.ID)
	}
	_, err = ParseToken(token, key+"x", sc)
	if err == nil {
		t.Error("invalid key but no return error")
	}
}

func TestRSAToken(t *testing.T) {
	key, _ := os.ReadFile("../assets/test_rsa.key")
	token, err := CreateTokenRSA(jwt.RegisteredClaims{
		ID: "123456",
	}, string(key))
	if err != nil {
		t.Error(err)
	}
	pub, _ := os.ReadFile("../assets/test_rsa.pem")
	sc := &jwt.RegisteredClaims{}
	_, err = ParseTokenRSA(token, string(pub), sc)
	if err != nil {
		t.Error(err)
	}
	if string(sc.ID) != "123456" {
		t.Error(sc.ID)
	}
	pub2, _ := os.ReadFile("../assets/test_rsa2.pem")
	_, err = ParseTokenRSA(token, string(pub2), sc)
	if err == nil {
		t.Error("invalid key but no return error")
	}
}
