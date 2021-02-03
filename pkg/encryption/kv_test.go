package encryption

import (
	"bytes"
	"context"
	"testing"

	"code.nfsmith.ca/nsmith/talaria/pkg/kv"
)

func TestBasics(t *testing.T) {
	k := kv.NewInMemory()
	k = &KVMiddleware{
		Next: k,
	}
	priv, pub, err := NewKeyPair()
	if err != nil {
		t.Fatal("Failed to generate keys")
	}

	value := []byte(`value`)
	ctx := WithKey(context.Background(), pub)

	err = k.Put(ctx, "key", value)
	if err != nil {
		t.Error("Failed to put")
	}

	ctx = WithKey(ctx, priv)
	returned, err := k.Get(ctx, "key")
	if err != nil {
		t.Error("Failed to get")
	}

	if equal := bytes.Compare(returned, value); equal != 0 {
		t.Error("Input doesn't match output")
	}

	err = k.Delete(ctx, "key")
	if err != nil {
		t.Error("Failed to delete key")
	}

	err = k.Delete(ctx, "key")
	if err != kv.ErrorNotFound {
		t.Error("Should fail on delete with no found")
	}
}
