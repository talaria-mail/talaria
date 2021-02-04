package mta

import (
	"net/mail"

	tmail "code.nfsmith.ca/nsmith/talaria/pkg/mail"
)

// Sender can send mail from one address to another
type Sender interface {
	Send(from mail.Address, to mail.Address, msg tmail.Message) error
}

// SenderFunc satisfies the Sender interface in a function
//
// SenderFunc is a convenience for writing Sender middleware (e.g DKIM etc)
type SenderFunc func(from mail.Address, to mail.Address, msg tmail.Message) error

func (f SenderFunc) Send(from mail.Address, to mail.Address, msg tmail.Message) error {
	return f(from, to, msg)
}
