package store

import (
	"context"
	"errors"

	"go.etcd.io/bbolt"
)

// Store is a message store. This includes all mailboxes for all users
type Store struct {
	DB *bbolt.DB

	ctx    context.Context
	cancel context.CancelFunc
}

// Run starts the mail store and blocks until Shutdown is called or a fatal
// error occurs
func (s *Store) Run() error {
	return errors.New("store: not implimented")
}

// Shutdown gracefully kills a running mailstore
func (s *Store) Shutdown(error) {
	return
}
