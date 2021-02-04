package logging

import (
	"log"
	"net/mail"

	tmail "code.nfsmith.ca/nsmith/talaria/pkg/mail"
	"code.nfsmith.ca/nsmith/talaria/pkg/mta"
)

// MTAMiddleware logs mta.Sender requests
func MTAMiddleware(next mta.Sender) mta.Sender {
	return mta.SenderFunc(func(from mail.Address, to mail.Address, msg tmail.Message) error {
		log.Printf("mta: from=%s to=%s subj=%s", from.Address, to.Address, msg.Header.Get("Subject"))
		return next.Send(from, to, msg)
	})
}
