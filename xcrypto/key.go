package xcrypto

import (
	"crypto/hmac"
	"crypto/rand"
	"io"

	"golang.org/x/crypto/hkdf"
	"golang.org/x/crypto/pbkdf2"
)

// GenExtendKey 使用 HKDF 生成扩展密钥
func GenExtendKey(key []byte, salt, info []byte, opts ...KeyOption) ([]byte, error) {
	c := applyKeyOption(opts...)
	h := hkdf.New(c.hash, key, salt, info)
	extendKey := make([]byte, c.keyLen)
	_, err := io.ReadFull(h, extendKey)
	if err != nil {
		return nil, err
	}
	return c.encode(extendKey), nil
}

// GenDeriveKey 使用 KDF 生成派生密钥
func GenDeriveKey(password, salt []byte, opts ...KeyOption) []byte {
	c := applyKeyOption(opts...)
	deriveKey := pbkdf2.Key(password, salt, c.iterations, c.keyLen, c.hash)
	return c.encode(deriveKey)
}

// GenKey 生成随机密钥
func GenKey(opts ...KeyOption) []byte {
	c := applyKeyOption(opts...)
	key := make([]byte, c.keyLen)
	_, _ = rand.Read(key)
	return c.encode(key)
}

// GenSalt 生成随机盐
func GenSalt(opts ...KeyOption) []byte {
	return GenKey(opts...)
}

// GenDEK 生成 DEK（Data Encryption Key，数据加密密钥）
func GenDEK(opts ...KeyOption) []byte {
	return GenKey(opts...)
}

// GenMasterKey 生成主密钥
func GenMasterKey(password, salt []byte, opts ...KeyOption) []byte {
	return GenDeriveKey(password, salt, opts...)
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
