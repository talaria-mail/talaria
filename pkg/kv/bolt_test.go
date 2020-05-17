package kv

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"testing"
)

func TestBoltBasics(t *testing.T) {
	dbfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal("Couldn't make file")
	}
	defer os.Remove(dbfile.Name())

	kv, err := NewBoltStore(dbfile.Name(), "bucket")
	if err != nil {
		t.Fatal("Couldn't make file")
	}

	value := []byte(`value`)
	ctx := context.Background()

	err = kv.Put(ctx, "key", value)
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
}

func TestBoltCancelation(t *testing.T) {
	dbfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal("Couldn't make file")
	}
	defer os.Remove(dbfile.Name())

	kv, err := NewBoltStore(dbfile.Name(), "bucket")
	if err != nil {
		t.Fatal("Couldn't make file")
	}

	value := []byte(`value`)
	ctx, cancel := context.WithCancel(context.Background())

	cancel()

	err = kv.Put(ctx, "key", value)
	if err != context.Canceled {
		t.Error("Should have failed to put on cancelled")
	}

	value, err = kv.Get(ctx, "key")
	if err != context.Canceled {
		t.Error("Should have failed to get on cancelled")
	}
}
