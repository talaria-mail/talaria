package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsmith5/talaria/pkg/auth"
	"github.com/nsmith5/talaria/pkg/kv"
	gauth "github.com/nsmith5/talaria/pkg/transport/grpc/auth"
	"github.com/nsmith5/talaria/pkg/users"

	"github.com/oklog/run"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
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

	var frontend http.Handler
	{
		mux := http.NewServeMux()
		mux.Handle("/", http.FileServer(AssetFile()))
		frontend = mux
	}

	var backend *grpc.Server
	{
		s := gauth.NewAuthServer(as)
		backend = grpc.NewServer()
		gauth.RegisterAuthServer(backend, s)
	}

	var g run.Group
	{
		server := &http.Server{
			Addr:    ":8080",
			Handler: frontend,
		}
		g.Add(func() error {
			log.Println("frontend: binding to 0.0.0.0:8080")
			return server.ListenAndServe()
		}, func(error) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			server.Shutdown(ctx)
		})
	}
	{
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 14586))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		g.Add(func() error {
			log.Print("backend: grpc binding to :14586")
			return backend.Serve(lis)
		}, func(error) {
			log.Print("backend: shutting down gracefully...")
			backend.GracefulStop()
		})
	}
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
