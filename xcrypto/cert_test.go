package xcrypto

import (
	"testing"
)

func TestCert(t *testing.T) {
	certText, prText, err := CreateCertToText(nil)
	if err != nil {
		t.Errorf("CreateCertToText error: %s", err)
	}
	t.Logf("cert: \n%s\nkey: \n%s", certText, prText)
	cert, err := TextToCert(certText)
	if err != nil {
		t.Errorf("TextToCert error: %s", err)
	}
	if cert.Subject.CommonName != "example.com" {
		t.Error("TextToCert common name is not example.com")
	}
}
