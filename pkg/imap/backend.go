package imap

import (
	"context"
	"errors"

	imap "github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend"
	"github.com/nsmith5/talaria/pkg/core"
)

type auth struct {
	core.AuthService
}

func (be *auth) Login(_ *imap.ConnInfo, username, password string) (backend.User, error) {
	ctx := context.Background()
	err := be.AuthService.Login(ctx, username, password)
	if err != nil {
		return nil, err
	}
	// TODO: Figure out user stuff?
	return nil, errors.New("Not implimented")
}

func newBackend() (backend.Backend, error) {
	return nil, errors.New("Not implimented")
}
