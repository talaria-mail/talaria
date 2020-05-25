package auth

import (
	"context"
	"testing"
)

func TestContext(t *testing.T) {
	token := "Hello World"
	ctx := context.Background()

	ctx = WithAuth(ctx, token)

	retrieved, ok := FromContext(ctx)
	if !ok {
		t.Error("Couldn't retrieve token")
	}

	if retrieved != token {
		t.Error("Input != output token")
	}
}
