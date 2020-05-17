package auth

import (
	"context"
	"errors"
)

// Authenticator handles login requests and returns JSON web tokens in exchange
// for credentials.
//
// Tokens are signed via elliptical curve crypto (for shorter tokens). This
// also means that public keys can be distributed to all middlewares and other
// services so that they can independently validate tokens
type Authenticator interface {
	Login(ctx context.Context, username, password string) (token string, err error)
}

// ErrorUnauthenticated signifies a failed login attempt
var ErrorUnauthenticated = errors.New("talaria/auth: unauthenticated")
