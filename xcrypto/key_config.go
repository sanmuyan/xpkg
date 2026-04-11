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
	*cryptoConfig
}

func newKeyConfig() *keyConfig {
	return &keyConfig{
		iterations:   Iterations,
		keyLen:       DefaultKeyLen,
		hash:         sha256.New,
		cryptoConfig: newCryptoConfig(),
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

func WithKeyIterations(iterations int) KeyOption {
	return func(c *keyConfig) {
		c.iterations = iterations
	}
}

func WithKeyLen(keyLen int) KeyOption {
	return func(c *keyConfig) {
		c.keyLen = keyLen
	}
}

func WithKeyHash(hash func() hash.Hash) KeyOption {
	return func(c *keyConfig) {
		c.hash = hash
	}
}

func WithKeyEncoder(encoder Encoder) KeyOption {
	return func(c *keyConfig) {
		c.encoder = encoder
	}
}
