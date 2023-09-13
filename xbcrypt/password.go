package xbcrypt

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

func IsPasswordComplexity(s string, minLength int, containsNumber, containsSymbols, containsLowercase, containsUppercase bool) bool {
	if !regexp.MustCompile(`^[\da-zA-Z!@#$%^&*]*$`).MatchString(s) {
		return false
	}
	if len(s) < minLength || len(s) > 32 {
		return false
	}
	if containsNumber {
		if !regexp.MustCompile(`\d`).MatchString(s) {
			return false
		}
	}
	if containsSymbols {
		if !regexp.MustCompile(`[!@#$%^&*]`).MatchString(s) {
			return false
		}
	}
	if containsLowercase {
		if !regexp.MustCompile(`[a-z]`).MatchString(s) {
			return false
		}
	}
	if containsUppercase {
		if !regexp.MustCompile(`[A-Z]`).MatchString(s) {
			return false
		}
	}
	return true
}

func CreatePassword(password string) (hashPassword string) {
	p := []byte(password)
	h, _ := bcrypt.GenerateFromPassword(p, bcrypt.MinCost)
	hashPassword = string(h)
	return hashPassword
}

func ComparePassword(hashPassword string, password string) bool {
	p := []byte(password)
	h := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(h, p)
	if err != nil {
		return false
	}
	return true
}
