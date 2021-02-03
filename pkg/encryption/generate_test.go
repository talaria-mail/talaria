package encryption

import (
	"bytes"
	"testing"

	"golang.org/x/crypto/curve25519"
)

func TestNewKeyPair(t *testing.T) {
	priv, pub, err := NewKeyPair()
	if err != nil {
		t.Error("Failed to create key pair")
	}

	pub2, err := curve25519.X25519(priv, curve25519.Basepoint)
	if err != nil {
		t.Fatal("Failed to generate second pub key")
	}

	if bytes.Compare(pub2, pub) != 0 {
		t.Error("pub key doesn't match expectation")
	}

	priv2, _, err := NewKeyPair()
	if err != nil {
		t.Fatal("Failed to generate a second key pair")
	}

	if bytes.Compare(priv2, priv) == 0 {
		t.Error("Failed to generate unique key pairs")
	}
}
