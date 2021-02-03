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

	_, err = Decrypt(make([]byte, 32), encrypted)
	if err == nil {
		t.Error("Decryption with an invalid private key shouldn't work")
	}

	_, err = Decrypt(make([]byte, 30), encrypted)
	if err == nil {
		t.Error("Decryption with an short private key shouldn't work")
	}

	_, err = Decrypt(make([]byte, 34), encrypted)
	if err == nil {
		t.Error("Decryption with an long private key shouldn't work")
	}

	_, err = Encrypt(make([]byte, 32), want)
	if err == nil {
		t.Error("Encryption with bad public key shouldn't wor")
	}
}
