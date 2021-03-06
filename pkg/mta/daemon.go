package mta

import (
	"context"

	"code.nfsmith.ca/nsmith/talaria/pkg/mail"
	"code.nfsmith.ca/nsmith/talaria/pkg/pubsub"
)

// Daemon pulls outbound messages from pubsub and sends them
type Daemon struct {
	PubSub pubsub.PubSub
	Sender Sender

	ctx    context.Context
	cancel context.CancelFunc
}

// Run starts the Daemon and blocks until an error occurs or Shutdown is called.
func (c *Daemon) Run() error {
	c.ctx, c.cancel = context.WithCancel(context.Background())

	events, errors := c.PubSub.Subscribe(c.ctx)
	for {
		select {
		// Shutdown
		case <-c.ctx.Done():
			return nil

		// Failures from PubSub
		case err := <-errors:
			return err

		// Happy path
		case event := <-events:
			switch msg := event.(type) {

			// Only subscribe to outbound message events
			case *mail.EventOutbound:

				// Loop the recepients and try to sent to each one
				for _, to := range msg.To {

					// DKIM and other concerns are assumed to be packaged as middleware on the
					// Sender
					err := c.Sender.Send(msg.From, to, msg.Message)
					if err != nil {
						// Create a error email from admin and send it to the
						// user to inform them of the delivery failure.
						errMsg := makeFailure(*msg, err)
						err = c.PubSub.Publish(context.Background(), &errMsg)
						if err != nil {
							// Failure to publish is assumed to be no recoverable. Stop the daemon
							return err
						}
					}
				}
			}
		}
	}
}

// Shutdown stops a running Daemon
func (c *Daemon) Shutdown(error) {
	c.cancel()
	return
}
