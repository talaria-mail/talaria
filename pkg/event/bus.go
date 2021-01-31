package event

import (
	"context"
	"sync"

	"code.nfsmith.ca/nsmith/talaria/pkg/talaria"
)

type bus struct {
	sync.Mutex

	subs map[*subscriber]struct{}
}

// NewBus returns a buffered event bus
func NewBus() talaria.Bus {
	return &bus{
		subs: map[*subscriber]struct{}{},
	}
}

func (b *bus) Publish(ctx context.Context, evt *talaria.Event) error {
	b.Lock()
	defer b.Unlock()
	for s := range b.subs {
		s.publish(evt)
	}
	return nil
}

func (b *bus) Subscribe(ctx context.Context) (<-chan *talaria.Event, <-chan error) {
	b.Lock()
	s := &subscriber{
		handler: make(chan *talaria.Event, 100),
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
