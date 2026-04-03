package xcrypto

import (
	"testing"
)

var key = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
var key2 = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxi"
var text = "123456"

func TestCFB(t *testing.T) {
	ciphertext, err := EncryptCFB(text, key)
	if err != nil {
		t.Error(err)
	}
	t.Logf("ciphertext: %s", ciphertext)
	if ciphertext == text {
		t.Error(ciphertext)
	}
	plaintext, err := DecryptCFB(ciphertext, key)
	if err != nil {
		t.Error(err)
	}
	t.Logf("plaintext: %s", plaintext)
	if plaintext != text {
		t.Error(plaintext)
	}
	_, err = DecryptCFB(ciphertext, key2)
	if err == nil {
		t.Errorf("invalid key but no return error")
	}

	type TestData struct {
		Ciphertext string
		T2         *struct {
			Ciphertext string
		}
		T3 *string
	}
	testData := &TestData{
		Ciphertext: ciphertext,
		T2: &struct {
			Ciphertext string
		}{
			Ciphertext: ciphertext,
		},
		T3: &ciphertext,
	}
	err = DecryptCFBToStruct(testData, key)
	if err != nil {
		t.Error(err)
	}
	if testData.Ciphertext != text || testData.T2.Ciphertext != text || *testData.T3 != text {
		t.Error(testData.Ciphertext, testData.T2.Ciphertext, *testData.T3)
	}
}

func TestGCM(t *testing.T) {
	ciphertext, err := EncryptGCM(text, key)
	if err != nil {
		t.Error(err)
	}
	t.Logf("ciphertext: %s", ciphertext)
	if ciphertext == text {
		t.Error(ciphertext)
	}
	plaintext, err := DecryptGCM(ciphertext, key)
	if err != nil {
		t.Error(err)
	}
	t.Logf("plaintext: %s", plaintext)
	if plaintext != text {
		t.Error(plaintext)
	}
	_, err = DecryptGCM(ciphertext, key2)
	if err == nil {
		t.Errorf("invalid key but no return error")
	}

	type TestData struct {
		Ciphertext string
		T2         *struct {
			Ciphertext string
		}
		T3 *string
	}
	testData := &TestData{
		Ciphertext: ciphertext,
		T2: &struct {
			Ciphertext string
		}{
			Ciphertext: ciphertext,
		},
		T3: &ciphertext,
	}
	err = DecryptGCMToStruct(testData, key)
	if err != nil {
		t.Error(err)
	}
	if testData.Ciphertext != text || testData.T2.Ciphertext != text || *testData.T3 != text {
		t.Error(testData.Ciphertext, testData.T2.Ciphertext, *testData.T3)
	}
}
