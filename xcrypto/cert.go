package xcrypto

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"time"
)

// CreateCert 创建一个 x509 证书
func CreateCert(template *x509.Certificate) ([]byte, *rsa.PrivateKey, error) {
	if template == nil {
		template = &x509.Certificate{
			SerialNumber: big.NewInt(1),
			Subject: pkix.Name{
				CommonName:   "example.com",
				Organization: []string{"Example Group"},
			},
			NotBefore:             time.Now(),
			NotAfter:              time.Now().Add(24 * time.Hour * 365 * 10),
			KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
			ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
			BasicConstraintsValid: true,
			IsCA:                  true}
	}
	pr, _, err := GenerateRSAKeyPair(2048)
	certDER, err := x509.CreateCertificate(rand.Reader, template, template, pr.Public(), pr)
	if err != nil {
		return nil, nil, err
	}
	return certDER, pr, nil
}

// CertToText 把证书转换成 PEM 文本
func CertToText(certDER []byte) ([]byte, error) {
	pemBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certDER,
	}
	var pemBuffer bytes.Buffer
	if err := pem.Encode(&pemBuffer, pemBlock); err != nil {
		return nil, err
	}
	return pemBuffer.Bytes(), nil
}

// TextToCert 把 PEM 文本转换成证书
func TextToCert(certText []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(certText)
	if block == nil {
		return nil, errors.New("failed to decode certificate")
	}
	return x509.ParseCertificate(block.Bytes)
}

// CreateCertToText 创建一个 x509 证书并转换为 PEM 格式
func CreateCertToText(template *x509.Certificate) ([]byte, []byte, error) {
	certDER, pr, err := CreateCert(template)
	if err != nil {
		return nil, nil, err
	}
	prText, err := RSAPrivateKeyToText(pr)
	if err != nil {
		return nil, nil, err
	}
	certText, err := CertToText(certDER)
	if err != nil {
		return nil, nil, err
	}
	return certText, prText, nil
}

// CreateCertToTLS 创建一个 x509 证书并转换为 TLS 配置
func CreateCertToTLS(template *x509.Certificate) (*tls.Config, error) {
	certDER, pr, err := CreateCert(template)
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		Certificates: []tls.Certificate{
			{
				Certificate: [][]byte{certDER},
				PrivateKey:  pr,
			},
		},
	}, nil
}
