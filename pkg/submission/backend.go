package submission

import (
	"context"
	"errors"

	"code.nfsmith.ca/nsmith/talaria/pkg/identity"
	"code.nfsmith.ca/nsmith/talaria/pkg/pubsub"
	smtp "github.com/emersion/go-smtp"
)

type backend struct {
	publisher pubsub.Publisher
	id        identity.Service
}

// Login handles a login command with username and password.
func (be *backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	ctx := context.Background()
	err := be.id.Authenticate(ctx, username, password)
	if err != nil {
		// Bury the error reason here so it doesn't leak to the client
		return nil, errors.New("Unauthenticated")
	}

	u, err := be.id.Get(ctx, username)
	if err != nil {
		return nil, errors.New("Unauthenticated")
	}

	return &session{msg: nil, publisher: be.publisher, user: *u}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (be *backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, smtp.ErrAuthRequired
}
