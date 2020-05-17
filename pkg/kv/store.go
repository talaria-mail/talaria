package kv

import (
	"context"
	"errors"
)

// ErrorNotFound is returned by a key value store that has no
// value for a key.
var ErrorNotFound = errors.New("talaria/kv: values not found for key")

// Store is a key-value store abstraction
type Store interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Put(ctx context.Context, key string, value []byte) error
}
