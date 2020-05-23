package api

import (
	"context"

	"github.com/nsmith5/talaria/pkg/auth"
	"github.com/nsmith5/talaria/pkg/servers/api/proto"
)

type authServer struct {
	auth.Authenticator
}

// newAuthServer created a GRPC server that handles authentication requests.
func newAuthServer(a auth.Authenticator) proto.AuthServer {
	return &authServer{a}
}

func (as *authServer) Login(ctx context.Context, request *proto.LoginRequest) (*proto.LoginResponse, error) {
	token, err := as.Authenticator.Login(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	return &proto.LoginResponse{Token: token}, nil
}
