package xcrypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// GenerateHmacSha256 生成 sha256 消息摘要，密钥长度不能小于32
func GenerateHmacSha256(message string, key string) (string, error) {
	if len(key) < 32 {
		return "", errors.New("key length invalid")
	}
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil)), nil
}

// GenerateHmacSha1 生成 sha1 消息摘要，密钥长度不能小于16
func GenerateHmacSha1(message string, key string) (string, error) {
	if len(key) < 16 {
		return "", errors.New("key length invalid")
	}
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil)), nil
}
