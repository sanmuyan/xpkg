package xbcrypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
	"time"
)

// IsPasswordComplexity 检查密码是否符合复杂度
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

// CreatePassword 创建密码
func CreatePassword(password string) (hashPassword string) {
	p := []byte(password)
	h, _ := bcrypt.GenerateFromPassword(p, bcrypt.MinCost)
	hashPassword = string(h)
	return hashPassword
}

// ComparePassword 比较密码
func ComparePassword(hashPassword string, password string) bool {
	p := []byte(password)
	h := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(h, p)
	if err != nil {
		return false
	}
	return true
}

// GetTOTPToken 生成基于时间的一次性密码 (TOTP)
func GetTOTPToken(secret string, interval uint8) (string, error) {
	secret = strings.ToUpper(strings.ReplaceAll(secret, " ", ""))

	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	counter := time.Now().Unix() / int64(interval)

	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, uint64(counter))

	// 创建 HMAC-SHA1 哈希
	hmacSha1 := hmac.New(sha1.New, key)
	hmacSha1.Write(counterBytes)
	hash := hmacSha1.Sum(nil)

	// 动态截取哈希值
	offset := hash[len(hash)-1] & 0x0F
	truncatedHash := binary.BigEndian.Uint32(hash[offset : offset+4])
	truncatedHash &= 0x7FFFFFFF // 只保留 31 位

	// 生成6位验证码
	code := truncatedHash % 1000000
	return fmt.Sprintf("%06d", code), nil
}
