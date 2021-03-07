package submission

import (
	"crypto/tls"

	"code.nfsmith.ca/nsmith/talaria/pkg/identity"
	"code.nfsmith.ca/nsmith/talaria/pkg/pubsub"
	smtp "github.com/emersion/go-smtp"
)

// Config configures a submission server
type Config struct {
	Addr   string
	Domain string
	TLS    tls.Config
}

func defaults(conf Config) Config {
	if conf.Addr == "" {
		conf.Addr = ":465"
	}

	if conf.Domain == "" {
		conf.Addr = "localhost"
	}

	return conf
}

// Server is a submission server
type Server struct {
	Config Config
	Pub    pubsub.Publisher
	ID     identity.Service

	s *smtp.Server
}

func (s *Server) Run() error {
	s.Config = defaults(s.Config)

	s.s = smtp.NewServer(&backend{publisher: s.Pub, id: s.ID})

	s.s.Domain = s.Config.Domain
	s.s.Addr = s.Config.Addr
	s.s.AllowInsecureAuth = false
	s.s.TLSConfig = &s.Config.TLS

	return s.s.ListenAndServeTLS()
}

func (s *Server) Shutdown(error) {
	s.s.Close()
	return
}
