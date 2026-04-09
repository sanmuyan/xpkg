package xcrypto

import (
	"encoding/base64"
	"testing"
)

var key = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
var text = "123456"

type TestDataDecryptToStruct struct {
	Ciphertext string
	T2         *struct {
		Ciphertext string
	}
	T3 *string
}

func TestCFB(t *testing.T) {
	ciphertext, err := EncryptCFB([]byte(text), []byte(key))
	if err != nil {
		t.Fatal(err)
	}
	plaintext, err := DecryptCFB(ciphertext, []byte(key))
	if err != nil {
		t.Fatal(err)
	}
	if string(plaintext) != text {
		t.Fatal("decrypt not expected")
	}
	t.Logf("plaintext: %s", plaintext)
	testDecryptStruct(t, ciphertext, DecryptCFB)
}

func TestGCM(t *testing.T) {
	ciphertext, err := EncryptGCM([]byte(text), []byte(key))
	if err != nil {
		t.Fatal(err)
	}
	plaintext, err := DecryptGCM(ciphertext, []byte(key))
	if err != nil {
		t.Fatal(err)
	}
	if string(plaintext) != text {
		t.Fatal("decrypt not expected")
	}
	t.Logf("plaintext: %s", plaintext)
	testDecryptStruct(t, ciphertext, DecryptGCM)
}

func testDecryptStruct(t *testing.T, ciphertext []byte, decryptFunc DecryptFunc) {
	ciphertextBase64 := base64.StdEncoding.EncodeToString(ciphertext)
	testData := &TestDataDecryptToStruct{
		Ciphertext: ciphertextBase64,
		T2: &struct {
			Ciphertext string
		}{
			Ciphertext: ciphertextBase64,
		},
		T3: &ciphertextBase64,
	}
	err := DecryptToStruct(testData, []byte(key), decryptFunc)
	if err != nil {
		t.Fatal(err)
	}
	if testData.Ciphertext != text || testData.T2.Ciphertext != text || *testData.T3 != text {
		t.Fatal(testData.Ciphertext, testData.T2.Ciphertext, *testData.T3)
	}
}
