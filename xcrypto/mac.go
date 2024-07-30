package xcrypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// GenerateHmacSha256 生成 hmac sha256
func GenerateHmacSha256(message string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
