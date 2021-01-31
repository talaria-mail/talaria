package submission

import (
	"code.nfsmith.ca/nsmith/talaria/pkg/pubsub"
	smtp "github.com/emersion/go-smtp"
)

type backend struct {
	publisher pubsub.Publisher
}

// Login handles a login command with username and password.
func (be *backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	// TODO: Impliment auth and push session token to session state
	return &session{msg: nil, publisher: be.publisher}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (be *backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, smtp.ErrAuthRequired
}
