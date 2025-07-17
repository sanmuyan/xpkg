package xcrypto

import (
	"testing"
)

func TestCert(t *testing.T) {
	cert, pr, err := CreateCert(nil)
	if err != nil {
		t.Error(err)
	}
	certText, prText, err := CertToText(cert, pr)
	if err != nil {
		t.Error(err)
	}
	t.Logf("cert: %s\nkey: %s", certText, prText)
}
