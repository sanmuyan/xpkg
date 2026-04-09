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
	}, []byte(key))
	if err != nil {
		t.Error(err)
	}
	t.Logf("token: %s", token)
	sc := &jwt.RegisteredClaims{}
	_, err = ParseToken(token, []byte(key), sc)
	if err != nil {
		t.Error(err)
	}
	if string(sc.ID) != "123456" {
		t.Error(sc.ID)
	}
	_, err = ParseToken(token, []byte(key+"x"), sc)
	if err == nil {
		t.Error("invalid key but no return error")
	}
}

func TestRSAToken(t *testing.T) {
	key, _ := os.ReadFile("../assets/test_cert.key")
	token, err := CreateTokenRSA(jwt.RegisteredClaims{
		ID: "123456",
	}, key)
	if err != nil {
		t.Error(err)
	}
	t.Logf("token: %s", token)
	pub, _ := os.ReadFile("../assets/test_cert.pem")
	sc := &jwt.RegisteredClaims{}
	_, err = ParseTokenRSA(token, pub, sc)
	if err != nil {
		t.Error(err)
	}
	if string(sc.ID) != "123456" {
		t.Error(sc.ID)
	}
	pub2, _ := os.ReadFile("../assets/test_rsa2.pem")
	_, err = ParseTokenRSA(token, pub2, sc)
	if err == nil {
		t.Error("invalid key but no return error")
	}
}
