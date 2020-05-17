package auth

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/nsmith5/talaria/pkg/kv"
	"github.com/nsmith5/talaria/pkg/users"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthenticator(t *testing.T) {
	password := "boo"
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user := users.User{
		Username:     "bob",
		PasswordHash: string(hash),
	}

	store := kv.NewMemStore()
	us := users.NewService(store)
	ctx := context.Background()
	us.Create(ctx, user)

	key, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		t.Error("Failed to generate key")
	}

	auth, err := NewAuthenticator(us, key)
	if err != nil {
		t.Fatal(err)
	}

	_, err = auth.Login(ctx, "bob", "boo")
	if err != nil {
		t.Error("Failed to login with correct credentials")
	}
}
