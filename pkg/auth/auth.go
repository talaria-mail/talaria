package auth

import (
	"context"

	"code.nfsmith.ca/nsmith/talaria/pkg/talaria"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Store talaria.UserStore
}

func (s *Service) Login(ctx context.Context, username, passwd string) (token string, err error) {
	user, err := s.Store.Get(ctx, username)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(user.PasswdHash, []byte(passwd))
	if err != nil {
		return
	}

	return
}
