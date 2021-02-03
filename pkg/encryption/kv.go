package encryption

import (
	"context"

	"code.nfsmith.ca/nsmith/talaria/pkg/kv"
)

type KVMiddleware struct {
	Next kv.Store
}

func (k *KVMiddleware) Get(ctx context.Context, key string) ([]byte, error) {
	encrypted, err := k.Next.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	privkey, err := KeyFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return Decrypt(privkey, encrypted)
}

func (k *KVMiddleware) Put(ctx context.Context, key string, value []byte) error {
	pubkey, err := KeyFromContext(ctx)
	if err != nil {
		return err
	}

	encrypted, err := Encrypt(pubkey, value)
	if err != nil {
		return err
	}

	return k.Next.Put(ctx, key, encrypted)
}

func (k *KVMiddleware) Delete(ctx context.Context, key string) error {
	encrypted, err := k.Next.Get(ctx, key)
	if err != nil {
		return err
	}

	privkey, err := KeyFromContext(ctx)
	if err != nil {
		return err
	}

	_, err = Decrypt(privkey, encrypted)
	if err != nil {
		return err
	}
	return k.Next.Delete(ctx, key)
}
