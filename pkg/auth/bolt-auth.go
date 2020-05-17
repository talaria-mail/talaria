package auth

import (
	"context"
	"crypto/ecdsa"

	"github.com/nsmith5/talaria/pkg/kv"
)

type kvAuth struct {
	kv.Store
	privateKey *ecdsa.PrivateKey
}

func NewKVAuth(store kv.Store, key *ecdsa.PrivateKey) Authenticator {
	return &kvAuth{store, key}
}

func (a *kvAuth) Login(ctx context.Context, username, password string) (token, error) {
	return "poop", nil
}
