package imap

import (
	"context"
	"strings"

	"github.com/emersion/go-imap/backend"
	"github.com/nsmith5/talaria/pkg/core"
)

type user struct {
	core.User
	core.MailboxStore
}

// Username returns this user's username.
func (u *user) Username() string {
	return u.Login
}

// ListMailboxes returns a list of mailboxes belonging to this user. If
// subscribed is set to true, only returns subscribed mailboxes.
func (u *user) ListMailboxes(subscribed bool) ([]backend.Mailbox, error) {
	ctx := context.Background()
	mbxs, err := u.List(ctx, u.ID, subscribed)
	if err != nil {
		return nil, err
	}

	var imbxs []backend.Mailbox
	for _, mbx := range mbxs {
		imbx := &mailbox{mbx, false}
		imbxs = append(imbxs, imbx)
	}
	return imbxs, nil
}

// GetMailbox returns a mailbox. If it doesn't exist, it returns
// ErrNoSuchMailbox.
func (u *user) GetMailbox(name string) (backend.Mailbox, error) {
	ctx := context.Background()
	mbx, err := u.Find(ctx, u.ID, name)
	if err != nil {
		// TODO: Handle other errors here perhaps
		return nil, backend.ErrNoSuchMailbox
	}
	return &mailbox{mbx, false}, nil
}

// CreateMailbox creates a new mailbox.
//
// If the mailbox already exists, an error must be returned. If the mailbox
// name is suffixed with the server's hierarchy separator character, this is a
// declaration that the client intends to create mailbox names under this name
// in the hierarchy.
//
// If the server's hierarchy separator character appears elsewhere in the
// name, the server SHOULD create any superior hierarchical names that are
// needed for the CREATE command to be successfully completed.  In other
// words, an attempt to create "foo/bar/zap" on a server in which "/" is the
// hierarchy separator character SHOULD create foo/ and foo/bar/ if they do
// not already exist.
//
// If a new mailbox is created with the same name as a mailbox which was
// deleted, its unique identifiers MUST be greater than any unique identifiers
// used in the previous incarnation of the mailbox UNLESS the new incarnation
// has a different unique identifier validity value.
func (u *user) CreateMailbox(name string) error {
	ctx := context.Background()
	parts := strings.Split(name, delimiter)
	for i := 0; i < len(parts); i++ {
		name = strings.Join(parts[:i], delimiter)
		err := u.Create(ctx, u.ID, name)
		if err != nil {
			return err
		}
	}
	return nil
}

// DeleteMailbox permanently remove the mailbox with the given name. It is an
// error to // attempt to delete INBOX or a mailbox name that does not exist.
//
// The DELETE command MUST NOT remove inferior hierarchical names. For
// example, if a mailbox "foo" has an inferior "foo.bar" (assuming "." is the
// hierarchy delimiter character), removing "foo" MUST NOT remove "foo.bar".
//
// The value of the highest-used unique identifier of the deleted mailbox MUST
// be preserved so that a new mailbox created with the same name will not
// reuse the identifiers of the former incarnation, UNLESS the new incarnation
// has a different unique identifier validity value.
func (u *user) DeleteMailbox(name string) error {
	ctx := context.Background()
	return u.Delete(ctx, u.ID, name)
}

// RenameMailbox changes the name of a mailbox. It is an error to attempt to
// rename from a mailbox name that does not exist or to a mailbox name that
// already exists.
//
// If the name has inferior hierarchical names, then the inferior hierarchical
// names MUST also be renamed.  For example, a rename of "foo" to "zap" will
// rename "foo/bar" (assuming "/" is the hierarchy delimiter character) to
// "zap/bar".
//
// If the server's hierarchy separator character appears in the name, the
// server SHOULD create any superior hierarchical names that are needed for
// the RENAME command to complete successfully.  In other words, an attempt to
// rename "foo/bar/zap" to baz/rag/zowie on a server in which "/" is the
// hierarchy separator character SHOULD create baz/ and baz/rag/ if they do
// not already exist.
//
// The value of the highest-used unique identifier of the old mailbox name
// MUST be preserved so that a new mailbox created with the same name will not
// reuse the identifiers of the former incarnation, UNLESS the new incarnation
// has a different unique identifier validity value.
//
// Renaming INBOX is permitted, and has special behavior.  It moves all
// messages in INBOX to a new mailbox with the given name, leaving INBOX
// empty.  If the server implementation supports inferior hierarchical names
// of INBOX, these are unaffected by a rename of INBOX.
func (u *user) RenameMailbox(existingName, newName string) error {
	ctx := context.Background()
	return u.Update(ctx, u.ID, existingName, newName)
}

// Logout is called when this User will no longer be used, likely because the
// client closed the connection.
func (u *user) Logout() error {
	return nil
}
