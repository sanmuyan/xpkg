package xcrypto

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"io"

	"golang.org/x/crypto/hkdf"
	"golang.org/x/crypto/pbkdf2"
)

const (
	Iterations    = 600000
	DefaultKeyLen = 32
)

// GenExtendKey 使用 HKDF 生成扩展密钥
func GenExtendKey(key []byte, salt, info []byte) ([]byte, error) {
	h := hkdf.New(sha256.New, key, salt, info)
	extendKey := make([]byte, DefaultKeyLen)
	_, err := io.ReadFull(h, extendKey)
	if err != nil {
		return nil, err
	}
	return extendKey, nil
}

// GenDeriveKey 使用 KDF 生成派生密钥
func GenDeriveKey(password, salt []byte) []byte {
	return pbkdf2.Key(password, salt, Iterations, DefaultKeyLen, sha256.New)
}

// GenSalt 生成随机盐
func GenSalt() []byte {
	salt := make([]byte, DefaultKeyLen)
	_, _ = rand.Read(salt)
	return salt
}

// GenDEK 生成 DEK（Data Encryption Key，数据加密密钥）
func GenDEK() []byte {
	dek := make([]byte, DefaultKeyLen)
	_, _ = rand.Read(dek)
	return dek
}

// GenMasterKey 生成主密钥
func GenMasterKey(password, salt []byte) []byte {
	return GenDeriveKey(password, salt)
}

// GenPasswordKDFHash 生成密码 KDF 哈希
func GenPasswordKDFHash(password, salt []byte) []byte {
	masterKey := GenMasterKey(password, salt)
	// 使用 masterKey 生成 masterKeyHash，这样服务端不用存储 masterKey
	masterKeyHash := GenMasterKey(masterKey, password)
	return masterKeyHash
}

// ComparePasswordKDFHash 对比密码 KDF 哈希
func ComparePasswordKDFHash(a, b []byte) bool {
	return hmac.Equal(a, b)
}
