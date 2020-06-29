package submission

import (
	"crypto/tls"
	"time"

	"github.com/emersion/go-smtp"
	"github.com/nsmith5/talaria/pkg/auth"
)

// Config holds submission server configuration details
type Config struct {
	Addr      string
	Auth      auth.Authenticator
	TLSConfig tls.Config
	Domain    string
}

// Server is a submission server.
type Server struct {
	s *smtp.Server
}

// New returns a new submission server
func New(cfg Config) Server {
	b := backend{
		auth: cfg.Auth,
	}

	s := smtp.NewServer(&b)
	s.Addr = cfg.Addr
	s.Domain = cfg.Domain
	s.TLSConfig = &cfg.TLSConfig

	s.LMTP = false
	s.MaxRecipients = 50
	s.MaxMessageBytes = 1024 * 2048
	s.MaxLineLength = 256
	s.AllowInsecureAuth = false
	s.Strict = true
	s.Debug = nil
	s.ErrorLog = nil
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second

	s.EnableSMTPUTF8 = true
	s.EnableREQUIRETLS = false
	s.AuthDisabled = false

	return Server{s}
}

// Run starts the server and returns an error if start up fails.
func (s *Server) Run() error {
	return s.s.ListenAndServeTLS()
}

// Shutdown gracefully stops the submission server.
func (s *Server) Shutdown(error) {
	s.s.Close()
	return
}
