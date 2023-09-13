package xbcrypt

import (
	"testing"
)

func TestPassword(t *testing.T) {
	if !IsPasswordComplexity("123AAaa#123", 8, true, true, true, true) {
		t.Error("invalid complexity")
	}
	if IsPasswordComplexity(" 123AAaa#123", 8, true, true, true, true) {
		t.Error("invalid complexity")
	}
	if IsPasswordComplexity("123aa#123", 8, true, true, true, true) {
		t.Error("invalid complexity")
	}
	hashPassword := CreatePassword("123AAaa#123")
	if !ComparePassword(hashPassword, "123AAaa#123") {
		t.Error("invalid password")
	}
	if ComparePassword(hashPassword, "123AAaa#123x") {
		t.Error("invalid password")
	}
}
