package xcrypto

import (
	"encoding/json"
	"testing"

	"github.com/sanmuyan/xpkg/xutil"
)

func TestVault(t *testing.T) {
	password := "my-password"
	newPassword := "new-password"
	msg := "hello world"
	masterKey := GenMasterKey([]byte(password), GenSalt())
	vault, err := CreateVault(masterKey, []byte(msg))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("vault: %s", xutil.RemoveError(json.Marshal(vault)))
	plaintext, err := DecryptVault(vault, masterKey)
	if err != nil {
		t.Fatal(err)
	}
	if string(plaintext) != msg {
		t.Fatalf("plaintext: %s != msg: %s", plaintext, msg)
	}
	t.Logf("plaintext: %s", plaintext)
	newMasterKey := GenMasterKey([]byte(newPassword), GenSalt())
	err = UpdateVaultKEK(vault, masterKey, newMasterKey, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("nwe vault: %s", xutil.RemoveError(json.Marshal(vault)))
	plaintext, err = DecryptVault(vault, newMasterKey)
	if err != nil {
		t.Fatal(err)
	}
	if string(plaintext) != msg {
		t.Fatalf("plaintext: %s != msg: %s", plaintext, msg)
	}
	t.Logf("new password plaintext: %s", plaintext)
}
