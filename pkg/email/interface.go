package email

import (
	"net/mail"
)

type Message struct {
	Metadata struct {
		ID    uint
		Box   string
		Flags uint
	}

	Header mail.Header
	Body   []byte
}

type Sender interface {
	Send(Message) bool
}

type Receiver interface {
	Receive(Message) bool
}

type FilterFunc func(Message) bool

type Repository interface {
	Receiver
	Select(FilterFunc) <-chan Message
	Delete(FilterFunc) uint
	Count(FilterFunc) uint
}
