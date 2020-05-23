package api

import (
	"context"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/nsmith5/talaria/pkg/auth"
	"github.com/nsmith5/talaria/pkg/servers/api/proto"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

// Config is an API server configuration.
type Config struct {
	Auth auth.Authenticator
	Addr string // Defaults to 0.0.0.0:8081
}

type Server struct {
	server *http.Server
}

// New returns a API GRPC server
func New(conf Config) Server {
	s := grpc.NewServer()
	as := newAuthServer(conf.Auth)
	proto.RegisterAuthServer(s, as)

	c := cors.AllowAll()
	handler := c.Handler(grpcweb.WrapServer(s))

	var addr string
	if conf.Addr == "" {
		addr = "0.0.0.0:8081"
	} else {
		addr = conf.Addr
	}

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return Server{server}
}

func (s Server) Run() error {
	return s.server.ListenAndServe()
}

func (s Server) Shutdown(error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s.server.Shutdown(ctx)
	return
}
