package talaria

import "context"

// AuthService handles creating sessions and their tokens
type AuthService interface {
	// Login creates a session and returns the token for it
	Login(ctx context.Context, username, password string) (token string, err error)

	// FetchSession information from token
	FetchSession(ctx context.Context, token string) (*Session, error)

	// RevokeSession removes a session even if it hasn't expired
	RevokeSession(ctx context.Context, token string) error

	// Storing token in context
	WithToken(ctx context.Context, token string) context.Context
	GetToken(ctx context.Context) (token string, err error)
}