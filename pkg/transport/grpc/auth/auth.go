package auth

import (
	"context"

	"github.com/nsmith5/talaria/pkg/auth"
)

//go:generate protoc --go_out=plugins=grpc:. auth.proto

type authServer struct {
	auth.Authenticator
}

// NewAuthServer created a GRPC server that handles authentication requests.
func NewAuthServer(a auth.Authenticator) AuthServer {
	return &authServer{a}
}

func (as *authServer) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	token, err := as.Authenticator.Login(ctx, request.Username, request.Password)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{Token: token}, nil
}
