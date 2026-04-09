package xmfa

import (
	"testing"
	"time"
)

func TestGetTOTPCode(t *testing.T) {
	token, err := GetTOTPCode("XGHRIBP4UHZ5TOIO", 30, time.Now().Unix())
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestValidateTOTPSecret(t *testing.T) {
	valid, err := ValidateTOTPCode("XGHRIBP4UHZ5TOIO", "123456", 30, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(valid)
}

func TestGenerateTOTPSecret(t *testing.T) {
	secret, err := GenerateTOTPSecret(16)
	if err != nil {
		t.Error(err)
	}
	t.Log(secret)
}
