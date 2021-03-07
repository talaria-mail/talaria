package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"code.nfsmith.ca/nsmith/talaria/pkg/logging"
	"code.nfsmith.ca/nsmith/talaria/pkg/mta"
	"code.nfsmith.ca/nsmith/talaria/pkg/pubsub"
	"code.nfsmith.ca/nsmith/talaria/pkg/submission"
	"github.com/oklog/run"
	"github.com/spf13/cobra"
)

func NewServeCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "serve",
		Short: "Run Talaria server daemon",
		Run:   RunServeCmd,
	}
	return cmd
}

func RunServeCmd(cmd *cobra.Command, args []string) {
	// TLS
	var tlsConf *tls.Config
	var err error
	if conf.TLS.Generate {
		tlsConf, err = TLSFromScratch(conf.Domain)
		if err != nil {
			panic(err)
		}
	} else {
		tlsConf, err = TLSFromFiles(conf.TLS.Cert, conf.TLS.Key)
		if err != nil {
			panic(err)
		}
	}

	// Pubsub event bus
	var ps pubsub.PubSub
	{
		ps = pubsub.NewPubSub()
		ps = &logging.PubSubMiddleware{Next: ps}
	}

	// Submission server
	var sub = submission.Server{
		Config: submission.Config{
			Addr:   fmt.Sprintf("0.0.0.0:%d", conf.Submission.Port),
			Domain: conf.Domain,
			TLS:    *tlsConf,
		},
		Pub: ps,
	}

	// MTA Sender
	var sender mta.Sender
	{
		sender = &mta.MailSender{
			Domain:  conf.Domain,
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

	err = g.Run()
	if err != nil {
		log.Println(err)
	}
}
