package xcrypto

type Encoder func(src []byte) []byte
type Decoder func(src []byte) ([]byte, error)

type cryptoConfig struct {
	encoder Encoder
	decoder Decoder
}

func newCryptoConfig() *cryptoConfig {
	return &cryptoConfig{}
}

type CryptoOption func(*cryptoConfig)

// WithEncryptEncoder 用于加密后数据编码
func WithEncryptEncoder(encoder Encoder) CryptoOption {
	return func(c *cryptoConfig) {
		c.encoder = encoder
	}
}

// WithDecryptDecoder 用于解密前数据解码
func WithDecryptDecoder(decoder Decoder) CryptoOption {
	return func(c *cryptoConfig) {
		c.decoder = decoder
	}
}

func applyCryptoOption(opts ...CryptoOption) *cryptoConfig {
	c := newCryptoConfig()
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *cryptoConfig) encode(src []byte) []byte {
	if c.encoder != nil {
		return c.encoder(src)
	}
	return src
}

func (c *cryptoConfig) decode(src []byte) ([]byte, error) {
	if c.decoder != nil {
		return c.decoder(src)
	}
	return src, nil
}
