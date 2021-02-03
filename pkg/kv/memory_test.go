package kv

import (
	"bytes"
	"context"
	"testing"
)

func TestBasics(t *testing.T) {
	kv := NewInMemory()
	value := []byte(`value`)
	ctx := context.Background()

	err := kv.Put(ctx, "key", value)
	if err != nil {
		t.Error("Failed to put")
	}

	returned, err := kv.Get(ctx, "key")
	if err != nil {
		t.Error("Failed to get")
	}

	if equal := bytes.Compare(returned, value); equal != 0 {
		t.Error("Input doesn't match output")
	}

	err = kv.Delete(ctx, "key")
	if err != nil {
		t.Error("Failed to delete key")
	}

	err = kv.Delete(ctx, "key")
	if err != ErrorNotFound {
		t.Error("Should fail on delete with no found")
	}
}

func TestCancelation(t *testing.T) {
	kv := NewInMemory()
	value := []byte(`value`)
	ctx, cancel := context.WithCancel(context.Background())

	cancel()

	err := kv.Put(ctx, "key", value)
	if err != context.Canceled {
		t.Error("Should have failed to put on cancelled")
	}

	value, err = kv.Get(ctx, "key")
	if err != context.Canceled {
		t.Error("Should have failed to get on cancelled")
	}

	err = kv.Delete(ctx, "key")
	if err != context.Canceled {
		t.Error("Should have failed with canceled")
	}
}
