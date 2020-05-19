package grpc

import (
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/nsmith5/talaria/pkg/auth"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

// NewServer returns a GRPC Web compatible GRPC server
func NewServer(authenticator auth.Authenticator) http.Handler {
	s := grpc.NewServer()
	as := NewAuthServer(authenticator)
	RegisterAuthServer(s, as)

	c := cors.AllowAll()
	return c.Handler(grpcweb.WrapServer(s))
}
