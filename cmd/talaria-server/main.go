package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"code.nfsmith.ca/nsmith/talaria/pkg/logging"
	"code.nfsmith.ca/nsmith/talaria/pkg/pubsub"
	"code.nfsmith.ca/nsmith/talaria/pkg/submission"
	"github.com/oklog/run"
)

type Config struct {
	submission submission.Config
}

func main() {
	var ps pubsub.PubSub
	{
		ps = pubsub.NewPubSub()
		ps = &logging.PubSubMiddleware{Next: ps}
	}

	conf := Config{
		submission: submission.Config{
			Addr:              ":6666",
			Domain:            "localhost",
			AllowInsecureAuth: true,
			Publisher:         ps,
		},
	}

	var subsServer submission.Server

	var g run.Group

	// Submission server
	g.Add(
		func() error {
			return subsServer.Start(conf.submission)
		}, func(err error) {
			fmt.Println(err)
			subsServer.Close()
			return
		},
	)

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
