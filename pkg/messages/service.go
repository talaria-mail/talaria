package messages

import (
	"context"
)

type Sender interface {
	Send(context.Context, Message) error
}

type Receiver interface {
	Receive(context.Context, Message) error
}

type Service interface {
	Sender
}
