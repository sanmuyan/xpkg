package xcrypto

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// IsPasswordComplexity 检查密码是否符合复杂度
func IsPasswordComplexity(s string, minLength int, minIncludeCase int) bool {
	if !regexp.MustCompile(`^[\da-zA-Z!@#$%^&*]*$`).MatchString(s) {
		return false
	}
	if len(s) < minLength || len(s) > 32 {
		return false
	}
	var ic int
	if regexp.MustCompile(`\d`).MatchString(s) {
		ic++
	}
	if regexp.MustCompile(`[!@#$%^&*]`).MatchString(s) {
		ic++
	}
	if regexp.MustCompile(`[a-z]`).MatchString(s) {
		ic++
	}
	if regexp.MustCompile(`[A-Z]`).MatchString(s) {
		ic++
	}
	if ic < minIncludeCase {
		return false
	}
	return true
}

// CreatePasswordHash 创建密码哈希值
func CreatePasswordHash(password []byte) (hashPassword []byte) {
	h, _ := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	return h
}

// ComparePassword 比较密码
func ComparePassword(hashPassword []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashPassword, password)
	if err != nil {
		return false
	}
	return true
}
