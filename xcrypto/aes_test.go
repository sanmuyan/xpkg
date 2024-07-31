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
	err = DecryptCFBToStruct(testData, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		t.Error(err)
	}
	if testData.Ciphertext != "123456" || testData.T2.Ciphertext != "123456" || *testData.T3 != "123456" {
		t.Error(testData.Ciphertext, testData.T2.Ciphertext, *testData.T3)
	}
}

func TestName2(t *testing.T) {
	x := "237b42a8-e485-452d-aa7b-0e0c670fe0fe"
	t.Log(EncryptCFB(x, "onoYKD9o59EhV9BFL1yu45PB702NR3bM"))
}
