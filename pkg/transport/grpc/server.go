package grpc

import (
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/nsmith5/talaria/pkg/auth"
	"google.golang.org/grpc"
)

// NewServer returns a GRPC Web compatible GRPC server
func NewServer(authenticator auth.Authenticator) http.Handler {
	s := grpc.NewServer()
	as := NewAuthServer(authenticator)
	RegisterAuthServer(s, as)

	return grpcweb.WrapServer(s)
}
