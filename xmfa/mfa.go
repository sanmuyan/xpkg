package xmfa

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

// GetTOTPToken 生成基于时间的一次性密码 (TOTP)
func GetTOTPToken(secret string, interval uint8, timestamp int64) (string, error) {
	secret = strings.ToUpper(strings.ReplaceAll(secret, " ", ""))

	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	counter := timestamp / int64(interval)

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

func GenerateTOTPSecret(length uint8) (string, error) {
	// 创建一个包含随机字节的数组
	secret := make([]byte, length)
	_, err := rand.Read(secret)
	if err != nil {
		return "", err
	}

	// 使用Base32编码
	base32Secret := base32.StdEncoding.EncodeToString(secret)

	// TOTP标准使用大写字母而没有填充字符"="
	return strings.TrimRight(base32Secret, "="), nil
}

// IsTOTPSecret 判断是否为合法的 TOTP 密钥
func IsTOTPSecret(secret string, length uint8) bool {
	if len(secret) != int(length) {
		return false
	}
	base32Chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	for _, char := range secret {
		if !strings.ContainsRune(base32Chars, char) {
			return false
		}
	}
	return true
}

// ValidateTOTPToken 验证 TOTP Token
func ValidateTOTPToken(secret string, userCode string, interval uint8, gracePeriod uint8) (bool, error) {
	currentTime := time.Now()
	count := 1
	if gracePeriod > 0 {
		count += int(gracePeriod)
	}
	for i := 0; i < count; i++ {
		expectedCode, err := GetTOTPToken(secret, interval, currentTime.Add(-time.Duration(i*int(interval))*time.Second).Unix())
		if err != nil {
			return false, err
		}
		if expectedCode == userCode {
			return true, nil
		}
	}

	return false, nil
}
