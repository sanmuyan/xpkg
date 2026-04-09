package xcrypto

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
)

// GenerateHmacSha256 生成 sha256 消息摘要，密钥长度必须为32
func GenerateHmacSha256(message []byte, key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("key length invalid")
	}
	h := hmac.New(sha256.New, key)
	h.Write(message)
	return h.Sum(nil), nil
}

// GenerateHmacSha1 生成 sha1 消息摘要，密钥长度必须为16
func GenerateHmacSha1(message []byte, key []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, errors.New("key length invalid")
	}
	h := hmac.New(sha1.New, []byte(key))
	h.Write(message)
	return h.Sum(nil), nil
}
