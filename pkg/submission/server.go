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
	Publisher         pubsub.Publisher
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
	s *smtp.Server
}

func (s *Server) Start(conf Config) error {
	conf = defaults(conf)

	s.s = smtp.NewServer(&backend{publisher: conf.Publisher})

	s.s.Domain = conf.Domain
	s.s.Addr = conf.Addr
	s.s.AllowInsecureAuth = conf.AllowInsecureAuth

	return s.s.ListenAndServe()
}

func (s *Server) Close() error {
	return s.s.Close()
}
