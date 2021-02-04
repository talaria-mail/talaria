package mail

import "net/mail"

// Message is an email message
type Message struct {
	Header mail.Header
	Body   []byte
}
