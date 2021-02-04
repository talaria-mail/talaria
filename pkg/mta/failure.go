package mta

import (
	"fmt"
	"time"

	"code.nfsmith.ca/nsmith/talaria/pkg/pubsub"
)

const errTemplate = `
`

func makeFailure(msg pubsub.EventOutbound, err error) pubsub.EventInbound {
	var r pubsub.EventInbound

	now := time.Now()

	// Pass along context from outbound message. This threading means encryption
	// context for the user will pass back with this failure message
	r.Context = msg.Context
	r.To = msg.From
	r.Message.Header = map[string][]string{
		"To":      {r.To.String()},
		"From":    {`admin@talaria.email`},
		"Subject": {`Failed to send message`},
		"Date":    {now.Format(`02 Jan 2006 15:04:05 -0700`)},
	}
	r.Message.Body = []byte(fmt.Sprintf("Error sending message %s", err))

	return r
}
