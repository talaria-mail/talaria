package user

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"

	"code.nfsmith.ca/nsmith/talaria/pkg/talaria"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/hkdf"
)

func deriveKey(passwd, salt []byte) (key []byte, err error) {
	// Derive Key encryption key
	kdf := hkdf.New(sha256.New, passwd, salt, nil)
	key = make([]byte, chacha20poly1305.KeySize)

	_, err = io.ReadFull(kdf, key)
	if err != nil {
		return nil, err
	}
	return
}

func encrypt(data, passwd, salt []byte) (encrypted []byte, err error) {
	// Key encryption key
	kek, err := deriveKey(passwd, salt)
	if err != nil {
		return nil, err
	}

	aead, err := chacha20poly1305.NewX(kek)
	if err != nil {
		return nil, err
	}

	// Select a random nonce, and leave capacity for the ciphertext.
	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(data)+aead.Overhead())
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err)
	}

	return aead.Seal(nonce, nonce, data, nil), nil
}

func decrypt(encrypted, passwd, salt []byte) (data []byte, err error) {
	// Key encryption key
	kek, err := deriveKey(passwd, salt)
	if err != nil {
		return nil, err
	}

	aead, err := chacha20poly1305.NewX(kek)
	if err != nil {
		return nil, err
	}

	if len(encrypted) < aead.NonceSize() {
		return nil, errors.New(`ciphertext too short (less then nonce)`)
	}

	// Split nonce and ciphertext.
	nonce, ciphertext := encrypted[:aead.NonceSize()], encrypted[aead.NonceSize():]

	// Decrypt the message and check it wasn't tampered with.
	data, err = aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return
}

// WithPassword sets a User password hash, salt and content key
//
// Should be used with new users for initial configuration. See ChangePassword
// for password updates.
func WithPassword(u talaria.User, passwd string) (*talaria.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), 15)
	if err != nil {
		return nil, err
	}
	u.PasswdHash = hash

	// Salt for HKDF of key encryption key
	salt := make([]byte, sha256.New().Size())
	_, err = rand.Read(salt)
	if err != nil {
		return nil, err
	}
	u.Salt = salt

	// New content encryption key (CEK)
	cek := make([]byte, chacha20poly1305.KeySize)
	_, err = rand.Read(cek)
	if err != nil {
		return nil, err
	}

	u.ContentKey, err = encrypt(cek, []byte(passwd), salt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// ChangePassword updates a User password hash, salt and content key using a new password
func ChangePassword(u talaria.User, old, new string) (*talaria.User, error) {
	// Check old password
	err := bcrypt.CompareHashAndPassword(u.PasswdHash, []byte(old))
	if err != nil {
		return nil, err
	}

	// Update password hash
	u.PasswdHash, err = bcrypt.GenerateFromPassword([]byte(new), 15)
	if err != nil {
		return nil, err
	}

	// New salt
	salt := make([]byte, sha256.New().Size())
	_, err = rand.Read(salt)
	if err != nil {
		return nil, err
	}
	u.Salt = salt

	// Content encryption key
	var cek = make([]byte, chacha20poly1305.KeySize)
	{
		cek, err = decrypt(u.ContentKey, []byte(old), u.Salt)
		if err != nil {
			return nil, err
		}
	}

	u.ContentKey, err = encrypt(cek, []byte(new), u.Salt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
