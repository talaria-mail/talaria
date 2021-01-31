package pubsub

import (
	"context"
	"sync"
)

type pubsub struct {
	sync.Mutex

	subs map[*subscriber]struct{}
}

// NewPubSub returns a buffered event bus
func NewPubSub() PubSub {
	return &pubsub{
		subs: map[*subscriber]struct{}{},
	}
}

func (b *pubsub) Publish(ctx context.Context, evt interface{}) error {
	b.Lock()
	defer b.Unlock()
	for s := range b.subs {
		s.publish(evt)
	}
	return nil
}

func (b *pubsub) Subscribe(ctx context.Context) (<-chan interface{}, <-chan error) {
	b.Lock()
	s := &subscriber{
		handler: make(chan interface{}, 100),
		quit:    make(chan struct{}),
	}
	b.subs[s] = struct{}{}
	b.Unlock()

	errc := make(chan error)
	go func() {
		defer close(errc)
		select {
		case <-ctx.Done():
			b.Lock()
			delete(b.subs, s)
			b.Unlock()
			s.close()
		}
	}()

	return s.handler, errc
}
