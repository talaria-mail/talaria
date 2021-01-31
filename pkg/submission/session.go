package submission

import (
	"context"
	"io"
	"io/ioutil"
	"net/mail"

	"code.nfsmith.ca/nsmith/talaria/pkg/pubsub"
	"code.nfsmith.ca/nsmith/talaria/pkg/talaria"
	smtp "github.com/emersion/go-smtp"
)

type session struct {
	msg       *talaria.OutboundMessage
	publisher pubsub.Publisher
}

func (s *session) Mail(from string, opts smtp.MailOptions) error {
	// Allocate new message
	s.msg = new(talaria.OutboundMessage)

	addr, err := mail.ParseAddress(from)
	if err != nil {
		return err
	}
	s.msg.From = *addr

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
