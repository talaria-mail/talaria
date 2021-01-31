package talaria

import "context"

// Bus is a multiple publisher, multiple subscriber bus that contains
// events for all message changes
type Bus interface {
	Publish(context.Context, *Event) error
	Subscribe(context.Context) (<-chan *Event, <-chan error)
}
