package xcrypto

import (
	"testing"
)

func TestCFB(t *testing.T) {
	ciphertext, err := EncryptCFB("123456", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		t.Error(err)
	}
	t.Logf("ciphertext: %s", ciphertext)
	if ciphertext == "123456" {
		t.Error(ciphertext)
	}
	plaintext, err := DecryptCFB(ciphertext, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		t.Error(err)
	}
	t.Logf("plaintext: %s", plaintext)
	if plaintext != "123456" {
		t.Error(plaintext)
	}
	_, err = DecryptCFB(ciphertext, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxi")
	if err == nil {
		t.Errorf("invalid key but no return error")
	}

	type TestData struct {
		Ciphertext string
		T2         *struct {
			Ciphertext string
		}
	}
	testData := &TestData{
		Ciphertext: ciphertext,
		T2: &struct {
			Ciphertext string
		}{
			Ciphertext: ciphertext,
		},
	}
	err = DecryptCFBToStruct(testData, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		t.Error(err)
	}
	if testData.Ciphertext != "123456" || testData.T2.Ciphertext != "123456" {
		t.Error(testData.Ciphertext, testData.T2.Ciphertext)
	}
}
