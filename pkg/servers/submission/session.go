package submission

import (
	"context"
	"io"
	"io/ioutil"
	"log"

	smtp "github.com/emersion/go-smtp"
)

// A Session is returned after successful login.
type session struct {
	ctx context.Context
}

func (s *session) Mail(from string, opts smtp.MailOptions) error {
	log.Println("Mail from:", from)
	return nil
}

func (s *session) Rcpt(to string) error {
	log.Println("Rcpt to:", to)
	return nil
}

func (s *session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		log.Println("Data:", string(b))
	}
	return nil
}

func (s *session) Reset() {}

func (s *session) Logout() error {
	return nil
}
