package submission

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/mail"

	"code.nfsmith.ca/nsmith/talaria/pkg/encryption"
	"code.nfsmith.ca/nsmith/talaria/pkg/identity"
	tmail "code.nfsmith.ca/nsmith/talaria/pkg/mail"
	"code.nfsmith.ca/nsmith/talaria/pkg/pubsub"
	smtp "github.com/emersion/go-smtp"
)

type session struct {
	msg       *tmail.EventOutbound
	publisher pubsub.Publisher
	user      identity.User
}

func (s *session) Mail(from string, opts smtp.MailOptions) error {
	// Allocate new message
	s.msg = &tmail.EventOutbound{
		Context: encryption.WithKey(context.Background(), s.user.ContentKey.Public),
	}

	// Parse address
	addr, err := mail.ParseAddress(from)
	if err != nil {
		return err
	}
	s.msg.From = *addr

	// Check if From address is allowed for this user
	err = tmail.AllowedAlias(addr.Address, s.user.EmailAliases)
	if err != nil {
		return errors.New(`invalid from address`)
	}
	return nil
}

func (s *session) Rcpt(to string) error {
	addr, err := mail.ParseAddress(to)
	if err != nil {
		return err
	}
	s.msg.To = append(s.msg.To, *addr)

	return nil
}

func (s *session) Data(r io.Reader) error {
	msg, err := mail.ReadMessage(r)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(msg.Body)
	if err != nil {
		return err
	}

	s.msg.Message.Header = msg.Header
	s.msg.Message.Body = body

	// Publish output message event
	err = s.publisher.Publish(context.Background(), s.msg)
	if err != nil {
		return err
	}

	s.msg = nil
	return nil
}

func (s *session) Reset() {
	s.msg = nil
}

func (s *session) Logout() error {
	return nil
}
