package core

import (
	"context"
)

type (
	User struct {
		ID    int
		Login string
	}
	UserStore interface {
		Find(context.Context, int) (*User, error)
		Create(context.Context, *User) error
		Update(context.Context, *User) error
		Delete(context.Context, *User) error
	}
)
