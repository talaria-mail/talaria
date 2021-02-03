package encryption

import (
	"crypto/rand"
	"errors"
	"io"

	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
)

// Encrypt uses a Curve25519 public key to encrypt `plain` using ECIES.
// Chacha20Poly1305 is used as the AEAD.
func Encrypt(pub []byte, plain []byte) (encrypted []byte, err error) {
	// Generate a random curve25519 key pair
	epriv, epub, err := NewKeyPair()
	if err != nil {
		return
	}

	// ECDH exchange to create a shared secret
	shared, err := curve25519.X25519(epriv, pub)
	if err != nil {
		return
	}

	// chacha20poly1305 aead encrypt plain with shared secret
	aead, err := chacha20poly1305.NewX(shared)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(plain)+aead.Overhead())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	encrypted = aead.Seal(nonce, nonce, plain, nil)

	// Prepend ephemeral public key to message
	encrypted = append(epub[:], encrypted...)
	return
}

// Decrypt uses a Curve25519 private key to decrypt `encrypted` using ECIES.
// Chacha20Poly1305 is used as the AEAD.
func Decrypt(priv []byte, encrypted []byte) (plain []byte, err error) {
	// Split ephemeral key out
	epub := encrypted[:32]
	encrypted = encrypted[32:]

	// ECDH to create a shared secret
	shared, err := curve25519.X25519(priv, epub)
	if err != nil {
		return
	}

	// chacha20poly1305 aead encrypt plain with shared secret
	aead, err := chacha20poly1305.NewX(shared)
	if err != nil {
		return nil, err
	}

	if len(encrypted) < aead.NonceSize() {
		return nil, errors.New("cipher text shorter than nonce")
	}

	// Split nonce and ciphertext.
	nonce, ciphertext := encrypted[:aead.NonceSize()], encrypted[aead.NonceSize():]

	// Decrypt the message and check it wasn't tampered with.
	return aead.Open(nil, nonce, ciphertext, nil)
}
