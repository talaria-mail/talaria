package logging

import (
	"context"
	"log"

	"code.nfsmith.ca/nsmith/talaria/pkg/pubsub"
)

// PubSubMiddleware logs pubsub events
type PubSubMiddleware struct {
	Next pubsub.PubSub
}

// Publish logs event and then publishes to Next
func (ps *PubSubMiddleware) Publish(ctx context.Context, evt interface{}) error {
	switch msg := evt.(type) {
	case *pubsub.EventOutbound:
		log.Printf("pubsub: method=Publish type=EventOutbound from=%s to=%s", msg.From, msg.To)
	case *pubsub.EventInbound:
		log.Printf("pubsub: method=Publish type=EventInbound to=%s", msg.To)
	}
	return ps.Next.Publish(ctx, evt)
}

// Subscribe logs event and then subscribes to Next
func (ps *PubSubMiddleware) Subscribe(ctx context.Context) (<-chan interface{}, <-chan error) {
	log.Println("pubsub: method=Subscribe")
	return ps.Next.Subscribe(ctx)
}
