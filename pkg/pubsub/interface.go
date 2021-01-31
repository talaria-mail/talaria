package pubsub

import "context"

type Publisher interface {
	Publish(context.Context, interface{}) error
}

type Subscriber interface {
	Subscribe(context.Context) (<-chan interface{}, <-chan error)
}

type PubSub interface {
	Publisher
	Subscriber
}
