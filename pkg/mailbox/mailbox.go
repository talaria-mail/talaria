package mailbox

import "net/mail"

type Message struct {
	ID     uint64
	Labels uint64
	Header mail.Header
	Body   []byte
}

type Mailbox interface {
	Insert(msg Message) error
	Find(id uint64) (Message, error)
	Update(id uint64, msg Message) error
	Delete(id uint64) error
}
