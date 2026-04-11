package xcrypto

import "testing"

func TestComparePasswordKDFHash(t *testing.T) {
	password := "my-password"
	password2 := "new-password"
	salt := GenSalt()
	masterKey := GenPasswordKDFHash([]byte(password), salt)
	expandMasterKeyHash := GenPasswordKDFHash(masterKey, []byte(password))

	masterKey2 := GenPasswordKDFHash([]byte(password), salt)
	expandMasterKeyHash2 := GenPasswordKDFHash(masterKey2, []byte(password))
	if !ComparePasswordKDFHash(expandMasterKeyHash, expandMasterKeyHash2) {
		t.Fatalf("not expected")
	}

	masterKey3 := GenPasswordKDFHash([]byte(password2), salt)
	expandMasterKeyHash3 := GenPasswordKDFHash(masterKey3, []byte(password2))
	if ComparePasswordKDFHash(expandMasterKeyHash, expandMasterKeyHash3) {
		t.Fatalf("not expected")
	}
}

func TestGenDeriveKey(t *testing.T) {
	GenDeriveKey([]byte("test"), []byte("test"), WithIterations(1000000))
}
