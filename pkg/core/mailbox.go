package core

import "context"

type (
	Mailbox interface {
		Name() string
		Count(ctx context.Context, qry MessageQuery) (int, error)
		Find(ctx context.Context, qry MessageQuery) ([]Message, error)
		First(ctx context.Context, qry MessageQuery) (*Message, error)
		Last(ctx context.Context, qry MessageQuery) (*Message, error)
	}

	MailboxStore interface {
		Create(ctx context.Context, userID int, name string) error
		Find(ctx context.Context, userID int, name string) (Mailbox, error)
		List(ctx context.Context, userID int, subscribed bool) ([]Mailbox, error)
		Update(ctx context.Context, userID int, oldname, newname string) error
		Delete(ctx context.Context, userID int, name string) error
	}
)
