package logging

import (
	"context"
	"log"

	"code.nfsmith.ca/nsmith/talaria/pkg/pubsub"
	"code.nfsmith.ca/nsmith/talaria/pkg/talaria"
)

type PubSubMiddleware struct {
	Next pubsub.PubSub
}

func (ps *PubSubMiddleware) Publish(ctx context.Context, evt interface{}) error {
	switch msg := evt.(type) {
	case *talaria.OutboundMessage:
		log.Printf("pubsub: method=Publish type=OutboundMessage from=%s to=%s", msg.From, msg.To)
	}
	return ps.Next.Publish(ctx, evt)
}

func (ps *PubSubMiddleware) Subscribe(ctx context.Context) (<-chan interface{}, <-chan error) {
	log.Println("pubsub: method=Subscribe")
	return ps.Next.Subscribe(ctx)
}
