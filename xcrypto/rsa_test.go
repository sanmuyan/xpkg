package xcrypto

import "testing"

func TestPKCSRSA(t *testing.T) {
	privateKey, publicKey, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Errorf("generate rsa key pair: %s", err)
	}
	privateKeyText, err := RSAPrivateKeyToText(privateKey)
	if err != nil {
		t.Errorf("to text rsa private key: %s", err)
	}
	publicKeyText, err := RSAPublicKeyToText(publicKey)
	if err != nil {
		t.Errorf("to text rsa public key: %s", err)
	}
	t.Logf("privateKey: \n%s", privateKeyText)
	t.Logf("publicKey: \n%s", publicKeyText)

	_privateKey, err := TextToRSAPrivateKey(privateKeyText)
	if err != nil {
		t.Errorf("from text rsa private key: %s", err)
	}
	_publicKey, err := TextToRSAPublicKey(publicKeyText)
	if err != nil {
		t.Errorf("from text public key: %s", err)
	}
	msg := "Hello World"
	ciphertext, err := EncryptPKCSRSA(msg, _publicKey)
	if err != nil {
		t.Errorf("encrypt rsa: %s", err)
	}
	t.Logf("ciphertext: %s", ciphertext)
	plaintext, err := DecryptPKCSRSA(ciphertext, _privateKey)
	if err != nil {
		t.Errorf("decrypt rsa: %s", err)
	}
	t.Logf("plaintext: %s", plaintext)
	if msg != plaintext {
		t.Errorf("msg: %s, plaintext: %s", msg, plaintext)
	}
}
