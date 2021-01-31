package talaria

import "context"

// Receiver
type Receiver interface {
	Receive(context.Context, Message) error
}
