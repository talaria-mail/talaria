package encryption

import (
	"crypto/rand"
	"io"

	"golang.org/x/crypto/curve25519"
)

// NewKeyPair creates a Curve25519 key pair
func NewKeyPair() (priv []byte, pub []byte, err error) {
	priv = make([]byte, 32)
	_, err = io.ReadFull(rand.Reader, priv)
	if err != nil {
		priv = nil
		pub = nil
		return
	}

	pub, err = curve25519.X25519(priv, curve25519.Basepoint)
	if err != nil {
		priv = nil
		pub = nil
		return
	}

	return priv, pub, nil
}
