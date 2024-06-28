package xmfa

import (
	"testing"
	"time"
)

func TestGetTOTPToken(t *testing.T) {
	token, err := GetTOTPToken("XGHRIBP4UHZ5TOIO", 30, time.Now().Unix())
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestValidateTOTPToken(t *testing.T) {
	valid, err := ValidateTOTPToken("XGHRIBP4UHZ5TOIO", "123456", 30, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(valid)
}
