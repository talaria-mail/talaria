package encryption

import (
	"crypto/rand"
	"errors"
	"io"

	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
)

func encrypt(pub [32]byte, plain []byte) (encrypted []byte, err error) {
	// Generate a random curve25519 key pair
	var epub, epriv [32]byte
	_, err = io.ReadFull(rand.Reader, epriv[:])
	if err != nil {
		return
	}
	curve25519.ScalarBaseMult(&epub, &epriv)

	// ECDH to create a shared secret
	var shared [32]byte
	curve25519.ScalarMult(&shared, &epriv, &pub)

	// chacha20poly1305 aead encrypt plain with shared secret
	aead, err := chacha20poly1305.NewX(shared[:])
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(plain)+aead.Overhead())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	encrypted = aead.Seal(nonce, nonce, plain, nil)

	// Append ephemeral public key to message
	encrypted = append(epub[:], encrypted...)
	return
}

func decrypt(priv [32]byte, encrypted []byte) (plain []byte, err error) {
	// Split ephemeral key out
	var epub [32]byte
	copy(epub[:], encrypted[:32])
	encrypted = encrypted[32:]

	// ECDH to create a shared secret
	var shared [32]byte
	curve25519.ScalarMult(&shared, &priv, &epub)

	// chacha20poly1305 aead encrypt plain with shared secret
	aead, err := chacha20poly1305.NewX(shared[:])
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
