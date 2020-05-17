package auth

import (
	"context"
	"crypto/ecdsa"

	"github.com/dgrijalva/jwt-go"
	"github.com/nsmith5/talaria/pkg/users"
)

type onlyAdmin struct {
	next users.Service
	key  ecdsa.PublicKey
}

// OnlyAdmin is a user service middleware that restricts user manipulation to
// admin users
func OnlyAdmin(next users.Service, key ecdsa.PublicKey) users.Service {
	return &onlyAdmin{next, key}
}

func (oa onlyAdmin) Create(ctx context.Context, user users.User) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		token, ok := FromContext(ctx)
		if !ok {
			return ErrorUnauthenticated
		}

		if !oa.verify(token) {
			return ErrorUnauthenticated
		}
		return oa.next.Create(ctx, user)
	}
}

func (oa onlyAdmin) Fetch(ctx context.Context, username string) (*users.User, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		token, ok := FromContext(ctx)
		if !ok {
			return nil, ErrorUnauthenticated
		}

		if !oa.verify(token) {
			return nil, ErrorUnauthenticated
		}
		return oa.next.Fetch(ctx, username)
	}
}

func (oa onlyAdmin) Update(ctx context.Context, user users.User) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		token, ok := FromContext(ctx)
		if !ok {
			return ErrorUnauthenticated
		}

		if !oa.verify(token) {
			return ErrorUnauthenticated
		}
		return oa.next.Update(ctx, user)
	}
}

func (oa onlyAdmin) Delete(ctx context.Context, username string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		token, ok := FromContext(ctx)
		if !ok {
			return ErrorUnauthenticated
		}

		if !oa.verify(token) {
			return ErrorUnauthenticated
		}
		return oa.next.Delete(ctx, username)
	}
}

func (oa onlyAdmin) verify(token string) bool {
	var claims jwt.MapClaims
	jwttoken, err := jwt.ParseWithClaims(
		token,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
				return nil, ErrorUnauthenticated
			}
			return oa.key, nil
		},
	)
	if err != nil || !jwttoken.Valid {
		return false
	}
	if !claims["admin"].(bool) {
		return false
	}
	return true
}
