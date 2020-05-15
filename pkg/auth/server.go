package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type authServer struct {
	jwtPrivatekey *rsa.PrivateKey
}

// NewServer creates a JWT powered auth token
func NewServer() AuthServer {
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}
	return &authServer{key}
}

func (as *authServer) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		jwt.MapClaims{
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
			"admin": true,
			"iss":   "auth.service",
			"iat":   time.Now().Unix(),
			"email": "root",
			"sub":   "root",
		},
	)

	tokenString, err := token.SignedString(as.jwtPrivatekey)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	return &LoginResponse{Token: tokenString}, nil
}
