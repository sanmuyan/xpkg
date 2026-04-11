package xcrypto

import (
	"encoding/base64"
	"testing"
)

var key = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
var text = "123456"

var postEncoder = WithPostEncoder(base64.StdEncoding.AppendEncode)
var preDecoder = WithPreDecoder(base64.StdEncoding.AppendDecode)

type TestDataDecryptToStruct struct {
	Ciphertext string
	T2         *struct {
		Ciphertext string
	}
	T3 *string
}

func TestCFB(t *testing.T) {
	ciphertext, err := EncryptCFB([]byte(text), []byte(key), postEncoder)
	if err != nil {
		t.Fatal(err)
	}
	plaintext, err := DecryptCFB(ciphertext, []byte(key), preDecoder)
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
	ciphertext, err := EncryptGCM([]byte(text), []byte(key), postEncoder)
	if err != nil {
		t.Fatal(err)
	}
	plaintext, err := DecryptGCM(ciphertext, []byte(key), preDecoder)
	if err != nil {
		t.Fatal(err)
	}
	if string(plaintext) != text {
		t.Fatal("decrypt not expected")
	}
	t.Logf("plaintext: %s", plaintext)
	testDecryptStruct(t, ciphertext, DecryptGCM)
}

func testDecryptStruct(t *testing.T, ciphertextBase64 []byte, decryptFunc DecryptFunc) {
	ciphertext := string(ciphertextBase64)
	testData := &TestDataDecryptToStruct{
		Ciphertext: ciphertext,
		T2: &struct {
			Ciphertext string
		}{
			Ciphertext: ciphertext,
		},
		T3: &ciphertext,
	}
	err := DecryptToStruct(testData, []byte(key), decryptFunc, preDecoder)
	if err != nil {
		t.Fatal(err)
	}
	if testData.Ciphertext != text || testData.T2.Ciphertext != text || *testData.T3 != text {
		t.Fatal(testData.Ciphertext, testData.T2.Ciphertext, *testData.T3)
	}
}
