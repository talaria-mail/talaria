package users

import (
	"context"
	"encoding/json"

	"github.com/nsmith5/talaria/pkg/kv"
)

type User struct {
	Username     string
	PasswordHash string
	IsAdmin      bool
	Email        string
	Aliases      []string
}

type Service interface {
	Create(context.Context, User) error
	Fetch(ctx context.Context, username string) (*User, error)
	Update(context.Context, User) error
	Delete(ctx context.Context, username string) error
}

type kvService struct {
	kv.Store
}

func NewService(store kv.Store) Service {
	return &kvService{store}
}

func (kv *kvService) Create(ctx context.Context, usr User) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		data, err := json.Marshal(usr)
		if err != nil {
			return err
		}
		return kv.Put(ctx, usr.Username, data)
	}
}

func (kv *kvService) Fetch(ctx context.Context, username string) (*User, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		data, err := kv.Get(ctx, username)
		if err != nil {
			return nil, err
		}

		var user User
		err = json.Unmarshal(data, &user)
		if err != nil {
			return nil, err
		}

		return &user, nil
	}
}

func (kv *kvService) Update(ctx context.Context, user User) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// Don't update if a user doesn't exist
		_, err := kv.Get(ctx, user.Username)
		if err != nil {
			return err
		}

		data, err := json.Marshal(user)
		if err != nil {
			return err
		}

		return kv.Put(ctx, user.Username, data)
	}
}

func (kv *kvService) Delete(ctx context.Context, username string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return kv.Store.Delete(ctx, username)
	}
}
