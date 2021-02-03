package encryption

import (
	"context"
	"errors"
)

type keyType struct{}

var k keyType

// WithKey adds a key (public or private) to context. This will be used for
// encryption should the context be passed the key value store encryption
// middleware
func WithKey(ctx context.Context, key []byte) context.Context {
	return context.WithValue(ctx, k, key)
}

func KeyFromContext(ctx context.Context) (key []byte, err error) {
	key, ok := ctx.Value(k).([]byte)
	if !ok {
		key = nil
		err = errors.New("Not key found")
		return
	}
	return
}
