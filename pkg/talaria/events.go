package talaria

import "net/mail"

// InboundMessage signals a message from the outside world to a users mailbox
type InboundMessage struct {
	Message Message
}

// OutboundMessage signals a message from a user to the outside world
type OutboundMessage struct {
	From    mail.Address
	To      []mail.Address
	Message Message
}
