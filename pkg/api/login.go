package api

import (
	"encoding/json"
	"net/http"

	"github.com/nsmith5/talaria/pkg/auth"
)

type loginRequest struct {
	Username string
	Password string
}

type loginResponse struct {
	Token string
}

func (a *api) HandleLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)

	var req loginRequest
	err := dec.Decode(&req)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	token, err := a.Login(r.Context(), req.Username, req.Password)
	if err == auth.ErrorUnauthenticated {
		http.Error(w, "Unauthenticated", http.StatusForbidden)
		return
	}
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
	}

	enc := json.NewEncoder(w)
	_ = enc.Encode(loginResponse{Token: token})
	return
}
