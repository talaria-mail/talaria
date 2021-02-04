package identity

import (
	"context"
	"errors"

	"code.nfsmith.ca/nsmith/talaria/pkg/kv"
)

type Service struct {
	KV kv.Store
}

func (s *Service) Authenticate(ctx context.Context, username, passwd string) error {
	return errors.New(`not implemented`)
}

func (s *Service) Get(ctx context.Context, username string) (*User, error) {
	return nil, errors.New(`not implemented`)
}
