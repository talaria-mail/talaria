package kv

import (
	"context"
)

type memKV map[string][]byte

func NewInMemory() Store {
	return make(memKV)
}

func (m memKV) Get(ctx context.Context, key string) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		value, ok := m[key]
		if !ok {
			return nil, ErrorNotFound
		}
		return value, nil
	}
}

func (m memKV) Put(ctx context.Context, key string, value []byte) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		m[key] = value
		return nil
	}
}

func (m memKV) Delete(ctx context.Context, key string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		_, ok := m[key]
		if !ok {
			return ErrorNotFound
		}
		delete(m, key)
		return nil
	}
}
