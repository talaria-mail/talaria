package encryption

import (
	"bytes"
	"testing"
)

func TestEncryption(t *testing.T) {
	// Generate key pair
	priv, pub, err := NewKeyPair()
	if err != nil {
		t.Fatal(`Failed to generate key pair`)
	}

	want := []byte(`this is my message`)

	encrypted, err := Encrypt(pub, want)
	if err != nil {
		t.Fatal("Failed to encrypt")
	}

	got, err := Decrypt(priv, encrypted)
	if err != nil {
		t.Fatal("Failed to decrypt")
	}

	if bytes.Compare(got, want) != 0 {
		t.Error("Failed to recover message")
	}
}
