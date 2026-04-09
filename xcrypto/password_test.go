package xcrypto

import (
	"testing"
)

func TestPassword(t *testing.T) {
	if !IsPasswordComplexity("123AAaa#123", 8, 4) {
		t.Error("invalid complexity")
	}
	if IsPasswordComplexity(" 123AAaa#123", 8, 4) {
		t.Error("invalid complexity")
	}
	if IsPasswordComplexity("123aa#123", 8, 4) {
		t.Error("invalid complexity")
	}
	hashPassword := CreatePasswordHash([]byte("123AAaa#123"))
	if !ComparePassword(hashPassword, []byte("123AAaa#123")) {
		t.Error("invalid password")
	}
	if ComparePassword(hashPassword, []byte("123AAaa#123x")) {
		t.Error("invalid password")
	}
}
