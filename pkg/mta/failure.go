package mta

import (
	"fmt"
	"time"

	"code.nfsmith.ca/nsmith/talaria/pkg/talaria"
)

const errTemplate = `
`

func makeFailure(msg talaria.OutboundMessage, err error) talaria.InboundMessage {
	var r talaria.InboundMessage

	now := time.Now()

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
