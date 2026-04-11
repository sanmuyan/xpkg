package xcrypto

import (
	"crypto/sha256"
	"hash"
)

const (
	// Iterations PBKDF2 迭代次数
	Iterations = 600000
	// DefaultKeyLen 默认密钥长度
	DefaultKeyLen = 32
)

type keyConfig struct {
	iterations int
	keyLen     int
	hash       func() hash.Hash
}

func newKeyConfig() *keyConfig {
	return &keyConfig{
		iterations: Iterations,
		keyLen:     DefaultKeyLen,
		hash:       sha256.New,
	}
}

type KeyOption func(*keyConfig)

func applyKeyOption(opts ...KeyOption) *keyConfig {
	c := newKeyConfig()
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// WithIterations 设置迭代次数
func WithIterations(iterations int) KeyOption {
	return func(c *keyConfig) {
		c.iterations = iterations
	}
}

// WithKeyLen 设置密钥长度
func WithKeyLen(keyLen int) KeyOption {
	return func(c *keyConfig) {
		c.keyLen = keyLen
	}
}

// WithHash 设置哈希函数
func WithHash(hash func() hash.Hash) KeyOption {
	return func(c *keyConfig) {
		c.hash = hash
	}
}
