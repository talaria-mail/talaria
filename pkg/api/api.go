package api

import (
	"net/http"

	"github.com/nsmith5/talaria/pkg/auth"
)

type api struct {
	*http.ServeMux
	auth.Authenticator
}

// New Talaria API server
func New(authenticator auth.Authenticator) http.Handler {
	a := &api{
		ServeMux:      http.NewServeMux(),
		Authenticator: authenticator,
	}
	a.HandleFunc("/login", a.HandleLogin)
	return a
}
