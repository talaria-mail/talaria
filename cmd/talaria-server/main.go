package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"code.nfsmith.ca/nsmith/talaria/pkg/logging"
	"code.nfsmith.ca/nsmith/talaria/pkg/mta"
	"code.nfsmith.ca/nsmith/talaria/pkg/pubsub"
	"code.nfsmith.ca/nsmith/talaria/pkg/submission"
	"github.com/oklog/run"
)

type Config struct {
	submission submission.Config
}

func main() {
	conf := Config{
		submission: submission.Config{
			Addr:              ":6666",
			Domain:            "localhost",
			AllowInsecureAuth: true,
		},
	}

	// Pubsub event bus
	var ps pubsub.PubSub
	{
		ps = pubsub.NewPubSub()
		ps = &logging.PubSubMiddleware{Next: ps}
	}

	// Submission server
	var sub = submission.Server{
		Config: conf.submission,
		Pub:    ps,
	}

	// MTA Sender
	var sender mta.Sender
	{
		sender = &mta.MailSender{
			Domain:  "localhost",
			Timeout: 10 * time.Second,
		}
		sender = logging.MTAMiddleware(sender)
	}

	var daemon = &mta.Daemon{
		PubSub: ps,
		Sender: sender,
	}

	// Error group orchestrates all of our processes together
	var g run.Group

	g.Add(sub.Run, sub.Shutdown)
	g.Add(daemon.Run, daemon.Shutdown)

	// Signal handler
	{
		ctx, cancel := context.WithCancel(context.Background())
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-ctx.Done():
				return ctx.Err()
			}
		}, func(error) {
			cancel()
		})
	}

	err := g.Run()
	if err != nil {
		fmt.Println(err)
	}
}
