package submission

import (
	"context"
	"io"
	"io/ioutil"

	smtp "github.com/emersion/go-smtp"
	"github.com/nsmith5/talaria/pkg/auth"
)

// A Session is returned after successful login.
type session struct {
	ctx  context.Context
	from *string
	to   []string
	data *string
}

func (s *session) Mail(from string, opts smtp.MailOptions) error {
	select {
	case <-s.ctx.Done():
		return s.ctx.Err()
	default:
		s.from = &from
		return nil
	}
}

func (s *session) Rcpt(to string) error {
	select {
	case <-s.ctx.Done():
		return s.ctx.Err()
	default:
		s.to = append(s.to, to)
		return nil
	}
}

func (s *session) Data(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	data := string(b)
	s.data = &data

	return nil
}

func (s *session) Reset() {
	s.from = nil
	s.to = nil
	s.data = nil
}

func (s *session) Logout() error {
	// Set JWT to empty
	s.ctx = auth.WithAuth(s.ctx, "")

	// Cancel the context
	ctx, cancel := context.WithCancel(s.ctx)
	cancel()
	s.ctx = ctx
	return nil
}
