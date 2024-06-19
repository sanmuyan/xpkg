package xcrypto

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// GenerateRSAKeyPair 生成 RSA 公钥私钥
func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

// RSAPrivateKeyToText RSA 私钥转成文本
func RSAPrivateKeyToText(key *rsa.PrivateKey) ([]byte, error) {
	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	bytBuff := bytes.NewBuffer([]byte{})
	err := pem.Encode(bytBuff, privateKey)
	if err != nil {
		return nil, err
	}
	return bytBuff.Bytes(), nil
}

// RSAPublicKeyToText RSA 公钥转成文本
func RSAPublicKeyToText(key *rsa.PublicKey) ([]byte, error) {
	asn1Bytes, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return nil, err
	}
	publicKey := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: asn1Bytes,
	}
	bytBuff := bytes.NewBuffer([]byte{})
	err = pem.Encode(bytBuff, publicKey)
	if err != nil {
		return nil, err
	}
	return bytBuff.Bytes(), nil
}

// TextToRSAPrivateKey RSA 私钥文本转为对象
func TextToRSAPrivateKey(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("private key error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// TextToRSAPublicKey RSA 公钥文本转为对象
func TextToRSAPublicKey(key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("public key error")
	}
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey, nil
}

// EncryptPKCSRSA PKCS RSA 加密
func EncryptPKCSRSA(plaintext string, key *rsa.PublicKey) (string, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, key, []byte(plaintext))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptPKCSRSA PKCS RSA 解密
func DecryptPKCSRSA(base64Ciphertext string, key *rsa.PrivateKey) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(base64Ciphertext)
	if err != nil {
		return "", err
	}
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, key, ciphertext)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
