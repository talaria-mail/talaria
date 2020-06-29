package submission

import (
	"context"

	smtp "github.com/emersion/go-smtp"
	"github.com/nsmith5/talaria/pkg/auth"
)

// The Backend implements SMTP server methods.
type backend struct {
	auth auth.Authenticator
}

// Login handles a login command with username and password.
func (b *backend) Login(state *smtp.ConnectionState, username, password string) (smtp.Session, error) {
	ctx := context.Background()
	token, err := b.auth.Login(ctx, username, password)
	if err != nil {
		return nil, &smtp.SMTPError{
			Code:         535,
			EnhancedCode: smtp.EnhancedCode{5, 3, 5},
			Message:      "Authentication failed",
		}
	}
	ctx = auth.WithAuth(ctx, token)
	return &session{ctx}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (b *backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	return nil, smtp.ErrAuthRequired
}
