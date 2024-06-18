package xmfa

import "testing"

func TestGetTOTPToken(t *testing.T) {
	token, err := GetTOTPToken("XGHRIBP4UHZ5TOIO", 30)
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
