package kv

import "context"

type prefixed struct {
	next   Store
	prefix string
}

// WithPrefix prepends a prefix to all keys passed to the underlying key
// value store.
func WithPrefix(s Store, prefix string) Store {
	return &prefixed{
		next:   s,
		prefix: prefix,
	}
}

func (p *prefixed) Get(ctx context.Context, key string) ([]byte, error) {
	return p.next.Get(ctx, p.prefix+key)
}

func (p *prefixed) Put(ctx context.Context, key string, value []byte) error {
	return p.next.Put(ctx, p.prefix+key, value)
}

func (p *prefixed) Delete(ctx context.Context, key string) error {
	return p.next.Delete(ctx, p.prefix+key)
}
