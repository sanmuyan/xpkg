package xcrypto

import "testing"

func TestGenerateRandomString(t *testing.T) {
	s := GenerateRandomString(64, true, true, true, true)
	t.Log(s)
}
