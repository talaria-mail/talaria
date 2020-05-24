package api

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/nsmith5/talaria/pkg/servers/api/proto"
	"github.com/nsmith5/talaria/pkg/users"
)

type usersServer struct {
	users.Service
}

// newUsersServer creates a GRPC server that handlers user CRUD requests
func newUserServer(srv users.Service) proto.UserServiceServer {
	return &usersServer{srv}
}

func (us *usersServer) Create(ctx context.Context, request *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	_, err := us.Service.Fetch(ctx, request.User.Username)
	if err == nil {
		return nil, errors.New("talaria/users: User already exists")
	}

	u := users.User{
		Username: request.User.Username,
		IsAdmin:  request.User.Admin,
		Email:    request.User.Email,
		Aliases:  request.User.Aliases,
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.User.Password), 13)
	if err != nil {
		return nil, err
	}

	u.PasswordHash = string(hash)

	err = us.Service.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return &proto.CreateUserResponse{}, nil
}

func (us *usersServer) Fetch(ctx context.Context, request *proto.FetchUserRequest) (*proto.FetchUserResponse, error) {
	u, err := us.Service.Fetch(ctx, request.Username)
	if err != nil {
		return nil, err
	}

	return &proto.FetchUserResponse{
		User: &proto.User{
			Username: u.Username,
			Password: ``,
			Admin:    u.IsAdmin,
			Email:    u.Email,
			Aliases:  u.Aliases,
		},
	}, nil
}

func (us *usersServer) List(ctx context.Context, request *proto.ListUsersRequest) (*proto.ListUsersResponse, error) {
	return nil, nil
}

func (us *usersServer) Update(ctx context.Context, request *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	return nil, nil
}

func (us *usersServer) Delete(ctx context.Context, request *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	return nil, nil
}
