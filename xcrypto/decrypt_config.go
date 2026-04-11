package xcrypto

type EncryptEncoder func(dst, src []byte) []byte
type EncryptDecoder func(dst, src []byte) ([]byte, error)

type encryptConfig struct {
	postEncoder EncryptEncoder
	preDecoder  EncryptDecoder
}

func newEncryptConfig() *encryptConfig {
	return &encryptConfig{}
}

type EncryptOption func(*encryptConfig)

func WithPostEncoder(encoder EncryptEncoder) EncryptOption {
	return func(c *encryptConfig) {
		c.postEncoder = encoder
	}
}

func WithPreDecoder(decoder EncryptDecoder) EncryptOption {
	return func(c *encryptConfig) {
		c.preDecoder = decoder
	}
}

func applyEncryptOption(opts ...EncryptOption) *encryptConfig {
	c := newEncryptConfig()
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func encryptEncode(src []byte, c *encryptConfig) []byte {
	if c.postEncoder != nil {
		return c.postEncoder(nil, src)
	}
	return src
}

func decryptDecode(src []byte, c *encryptConfig) ([]byte, error) {
	if c.preDecoder != nil {
		return c.preDecoder(nil, src)
	}
	return src, nil
}
