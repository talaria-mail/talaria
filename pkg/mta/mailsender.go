package mta

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"strings"
	"time"

	"code.nfsmith.ca/nsmith/talaria/pkg/talaria"
)

// MailSender sends mails and satifies the Sender interface
type MailSender struct {
	Resolver net.Resolver
	Domain   string
	Timeout  time.Duration
}

func (s *MailSender) Send(from mail.Address, to mail.Address, msg talaria.Message) error {
	d, err := domain(to.Address)
	if err != nil {
		return err
	}

	mxs, err := s.Resolver.LookupMX(context.Background(), d)
	if err != nil {
		return err
	}
	if len(mxs) == 0 {
		return fmt.Errorf("No mail exchanges found for domain %s", d)
	}

	var errs = make(map[string]error)
	for _, mx := range mxs {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:25", mx.Host), s.Timeout)
		if err != nil {
			return err
		}

		c, err := smtp.NewClient(conn, s.Domain)
		if err != nil {
			// Failure is _ok_, we try the next mail exchange
			errs[mx.Host] = err
			continue
		}

		err = c.Mail(from.String())
		if err != nil {
			c.Quit()
			c.Close()
			errs[mx.Host] = err
			continue
		}

		err = c.Rcpt(to.String())
		if err != nil {
			c.Quit()
			c.Close()
			errs[mx.Host] = err
			continue
		}

		wc, err := c.Data()
		if err != nil {
			c.Quit()
			c.Close()
			errs[mx.Host] = err
			continue
		}
		n, err := wc.Write(msg.Body)
		if err != nil {
			c.Quit()
			c.Close()
			errs[mx.Host] = err
			continue
		}
		if n != len(msg.Body) {
			c.Quit()
			c.Close()
			errs[mx.Host] = errors.New("Partial write of message body")
			continue
		}
		err = wc.Close()
		if err != nil {
			c.Quit()
			c.Close()
			errs[mx.Host] = err
			continue
		}

		// Success!
		c.Quit()
		c.Close()
		return nil
	}

	return fmt.Errorf("%v", errs)
}

func domain(address string) (string, error) {
	parts := strings.Split(address, `@`)
	if len(parts) != 2 {
		return "", errors.New("Invalid address")
	}
	return parts[1], nil
}
