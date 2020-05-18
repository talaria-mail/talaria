package auth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nsmith5/talaria/pkg/users"
	"golang.org/x/crypto/bcrypt"
)

// Authenticator handles login requests and returns JSON web tokens in
// exchange for credentials.
//
// Tokens are signed via elliptical curve crypto (for shorter tokens). This
// also means that public keys can be distributed to all middlewares and other
// services so that they can independently validate tokens
type Authenticator interface {
	Login(ctx context.Context, username, password string) (token string, err error)
}

// ErrorUnauthenticated signifies a failed login attempt
var ErrorUnauthenticated = errors.New("talaria/auth: unauthenticated")

type authenticator struct {
	us  users.Service
	key *ecdsa.PrivateKey
}

func NewAuthenticator(us users.Service, key *ecdsa.PrivateKey) (Authenticator, error) {
	if key == nil {
		return nil, errors.New("talaria/auth: nil private key")
	}
	return &authenticator{us, key}, nil
}

func (a *authenticator) Login(ctx context.Context, username, password string) (token string, err error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		user, err := a.us.Fetch(ctx, username)
		if err != nil {
			// It is important not to leak _why_ this failed (that the user
			// doesn't exist).
			return "", ErrorUnauthenticated
		}

		// Check password
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		if err != nil {
			return "", ErrorUnauthenticated
		}

		token := jwt.NewWithClaims(jwt.SigningMethodES512, jwt.MapClaims{
			// Standard claims
			"iss": "talaria.auth",
			"sub": user.Username,
			"aud": "talaria",
			"exp": time.Now().Add(time.Hour * 24).Unix(),
			"iat": time.Now().Unix(),
			"jti": 1,

			// Custom cliams
			"admin": user.IsAdmin,
		})

		tokenString, err := token.SignedString(a.key)
		if err != nil {
			return "", ErrorUnauthenticated
		}
		return tokenString, nil
	}
}
