package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsmith5/talaria/pkg/auth"
	"github.com/nsmith5/talaria/pkg/kv"
	"github.com/nsmith5/talaria/pkg/servers/api"
	"github.com/nsmith5/talaria/pkg/servers/web"
	"github.com/nsmith5/talaria/pkg/users"

	"github.com/oklog/run"
	"github.com/spf13/cobra"
)

//go:generate go-bindata -prefix "../../frontend/dist/" -fs ../../frontend/dist/...

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Talaria server",
	Run:   runServerCmd,
}

func runServerCmd(cmd *cobra.Command, args []string) {
	var store kv.Store = kv.NewMemStore()

	var (
		us users.Service
		as auth.Authenticator
	)
	{
		us = users.NewService(store)
		privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
		if err != nil {
			panic(err)
		}
		as, err = auth.NewAuthenticator(us, privateKey)
		if err != nil {
			panic(err)
		}
		us = auth.OnlyAdmin(us, privateKey.PublicKey)
	}

	var frontend web.Server
	{
		config := web.Config{
			Addr: "0.0.0.0:8080",
		}
		frontend = web.New(config)
	}

	var backend api.Server
	{
		config := api.Config{
			Auth: as,
			Addr: "0.0.0.0:8081",
		}
		backend = api.New(config)
	}

	var g run.Group
	g.Add(frontend.Run, frontend.Shutdown)
	g.Add(backend.Run, backend.Shutdown)
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

	fmt.Println(g.Run())
}
