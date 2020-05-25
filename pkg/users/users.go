package users

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"

	"golang.org/x/crypto/bcrypt"

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
	service := &kvService{store}
	ctx := context.Background()
	_, err := service.Fetch(ctx, `root`)
	if err == kv.ErrorNotFound {
		mustCreateRootUser(service)
	}

	return &kvService{store}
}

func mustCreateRootUser(srv Service) {
	var buf [32]byte
	_, err := io.ReadFull(rand.Reader, buf[:])
	if err != nil {
		panic(err)
	}
	password := base64.StdEncoding.EncodeToString(buf[:])
	log.Printf("talaria/users: welcome to talaria! root password is %s", password)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}

	err = srv.Create(
		context.Background(),
		User{Username: `root`, PasswordHash: string(hash), IsAdmin: true},
	)
	if err != nil {
		panic(err)
	}
	return
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
