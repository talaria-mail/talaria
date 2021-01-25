package core

import "context"

type (
	Mailbox interface {
		Name() string
	}
	MailboxStore interface {
		Create(ctx context.Context, userID int, name string) error
		Find(ctx context.Context, userID int, name string) (Mailbox, error)
		List(ctx context.Context, userID int, subscribed bool) ([]Mailbox, error)
		Update(ctx context.Context, userID int, oldname, newname string) error
		Delete(ctx context.Context, userID int, name string) error
	}
)
