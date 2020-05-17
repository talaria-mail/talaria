package auth

import "context"

type keyType struct{}

var key keyType

// WithAuth adds a JWT token to a context
func WithAuth(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, key, token)
}

// FromContext extracts authentication tokens from a context
func FromContext(ctx context.Context) (token string, ok bool) {
	token, ok = ctx.Value(key).(string)
	return
}
