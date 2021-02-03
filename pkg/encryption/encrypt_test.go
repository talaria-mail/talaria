package encryption

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"

	"golang.org/x/crypto/curve25519"
)

func TestEncryption(t *testing.T) {
	// Generate key pair
	var pub, priv [32]byte
	_, err := io.ReadFull(rand.Reader, priv[:])
	if err != nil {
		t.Fatal("Failed to create private key")
	}
	curve25519.ScalarBaseMult(&pub, &priv)

	want := []byte(`this is my message`)

	encrypted, err := encrypt(pub, want)
	if err != nil {
		t.Fatal("Failed to encrypt")
	}

	got, err := decrypt(priv, encrypted)
	if err != nil {
		t.Fatal("Failed to decrypt")
	}

	if bytes.Compare(got, want) != 0 {
		t.Error("Failed to recover message")
	}
}
