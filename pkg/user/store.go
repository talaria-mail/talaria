package user

import (
	"context"
	"errors"

	"code.nfsmith.ca/nsmith/talaria/pkg/talaria"
)

type store struct {
}

func NewStore() talaria.UserStore {
	return &store{}
}

func (s *store) Create(ctx context.Context, username string, password string) error {
	return errors.New("Not implimented")
}

func (s *store) Get(ctx context.Context, id []byte) (*talaria.User, error) {
	return nil, errors.New("Not implemented")
}

func (s *store) Update(ctx context.Context, user talaria.User) error {
	return errors.New("not implemented")
}

func (s *store) Delete(ctx context.Context, id []byte) error {
	return errors.New("not implemented")
}
