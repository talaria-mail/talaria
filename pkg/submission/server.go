package submission

import (
	"code.nfsmith.ca/nsmith/talaria/pkg/pubsub"
	smtp "github.com/emersion/go-smtp"
)

// Config configures a submission server
type Config struct {
	Addr              string
	Domain            string
	AllowInsecureAuth bool
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

	s *smtp.Server
}

func (s *Server) Run() error {
	s.Config = defaults(s.Config)

	s.s = smtp.NewServer(&backend{publisher: s.Pub})

	s.s.Domain = s.Config.Domain
	s.s.Addr = s.Config.Addr
	s.s.AllowInsecureAuth = s.Config.AllowInsecureAuth

	return s.s.ListenAndServe()
}

func (s *Server) Shutdown(error) {
	s.s.Close()
	return
}
