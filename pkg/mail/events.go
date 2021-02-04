package mail

import (
	"context"
	"net/mail"
)

// EventInbound signals a message from the outside world to a users mailbox
type EventInbound struct {
	Context context.Context
	To      mail.Address
	Message Message
}

// EventOutbound signals a message from a user to the outside world
type EventOutbound struct {
	Context context.Context
	From    mail.Address
	To      []mail.Address
	Message Message
}
