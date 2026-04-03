package xcrypto

import (
	"testing"
)

func TestCert(t *testing.T) {
	certText, prText, err := CreateCertToText(nil)
	if err != nil {
		t.Errorf("CreateCertToText error: %s", err)
	}
	t.Logf("cert: %s\nkey: %s", certText, prText)
	cert, err := TextToCert(certText)
	if err != nil {
		t.Errorf("TextToCert error: %s", err)
	}
	t.Logf("cert common name: %s", cert.Subject.CommonName)
}
