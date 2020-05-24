package api

import (
	"context"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/nsmith5/talaria/pkg/auth"
	"github.com/nsmith5/talaria/pkg/servers/api/proto"
	"github.com/nsmith5/talaria/pkg/users"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

// Config is an API server configuration.
type Config struct {
	Auth        auth.Authenticator
	UserService users.Service
	Addr        string // Defaults to 0.0.0.0:8081
}

// Server is an instance of a GRPC server
type Server struct {
	server *http.Server
}

// New returns a API GRPC server
func New(conf Config) Server {
	s := grpc.NewServer()
	as := newAuthServer(conf.Auth)
	proto.RegisterAuthServer(s, as)
	us := newUserServer(conf.UserService)
	proto.RegisterUserServiceServer(s, us)

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

// Run starts a GRPC server and exits on error
func (s Server) Run() error {
	return s.server.ListenAndServe()
}

// Shutdown gracefully shutsdown an API server
func (s Server) Shutdown(error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	s.server.Shutdown(ctx)
	return
}
