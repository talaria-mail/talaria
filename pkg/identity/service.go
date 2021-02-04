package identity

import (
	"bytes"
	"context"
	"crypto/subtle"
	"encoding/gob"
	"errors"

	"code.nfsmith.ca/nsmith/talaria/pkg/encryption"
	"code.nfsmith.ca/nsmith/talaria/pkg/kv"
	"golang.org/x/crypto/blake2b"
)

type Service struct {
	KV kv.Store
}

func (s *Service) Authenticate(ctx context.Context, username, passwd string) error {
	user, err := s.Get(ctx, username)
	if err != nil {
		return err
	}

	hash := encryption.HashPassword([]byte(passwd), user.Salt)
	if subtle.ConstantTimeCompare(hash, user.PasswordHash) != 1 {
		return errors.New(`Not authorized`)
	}
	return nil
}

func (s *Service) Get(ctx context.Context, username string) (*User, error) {
	id := blake2b.Sum256([]byte(username))
	blob, err := s.KV.Get(ctx, string(id[:]))
	if err != nil {
		return nil, err
	}

	var user User
	{
		dec := gob.NewDecoder(bytes.NewReader(blob))
		err = dec.Decode(&user)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}
