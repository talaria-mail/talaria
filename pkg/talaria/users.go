package talaria

import "context"

// User contains user metadata and authentication details
type User struct {
	// Login username
	Login string

	// Authentication and encryption meatadata
	PasswdHash []byte // bcrypt cost 15
	Salt       []byte // salt to derive key encryption key
	ContentKey []byte // encrypted content encryption key
}

// UserStore manages users
type UserStore interface {
	Create(ctx context.Context, username, password string) error
	Get(ctx context.Context, username string) (*User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, username string) error
}
