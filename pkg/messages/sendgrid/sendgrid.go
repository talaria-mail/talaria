package sendgrid

import (
	"context"
	"encoding/json"
	"os"

	"github.com/nsmith5/talaria/pkg/messages"
	"github.com/sendgrid/sendgrid-go"
)

type Addressee struct {
	Email string
	Name  *string
}

type Attachment struct {
	Content     []byte
	ContentID   string
	Disposition string
	Filename    string
	Type        string
}

type Personalization struct {
	To  []Addressee
	CC  []Addressee
	BCC []Addressee
}

type Request struct {
	Personalizations [1]Personalization
	From             Addressee
	Content          []struct {
		Type  string
		Value string
	}
	Attachments []Attachment
}

var (
	host     = `https://api.sendgrid.com`
	endpoint = `/v3/mail/send`
	apiKey   = os.Getenv("SENDGRID_API_KEY")
)

type sender struct{}

func NewSender() messages.Sender {
	return sender{}
}

func (s sender) Send(ctx context.Context, msg messages.Message) error {
	request := sendgrid.GetRequest(apiKey, endpoint, host)
	request.Method = "POST"

	req := toSendGridRequest(msg)
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	request.Body = data
	_, err = sendgrid.API(request)
	return err
}

func toSendGridRequest(msg messages.Message) Request {
	var (
		req     Request
		to      []Addressee
		cc      []Addressee
		bcc     []Addressee
		content []struct {
			Type  string
			Value string
		}
		attachments []Attachment
	)

	for i := 0; i < len(msg.To); i++ {
		to = append(to, Addressee{Email: msg.To[i]})
	}
	for i := 0; i < len(msg.CC); i++ {
		cc = append(cc, Addressee{Email: msg.CC[i]})
	}
	for i := 0; i < len(msg.BCC); i++ {
		bcc = append(bcc, Addressee{Email: msg.BCC[i]})
	}

	req.Personalizations[0].To = to
	req.Personalizations[0].CC = cc
	req.Personalizations[0].BCC = bcc

	req.From = Addressee{Email: msg.From}

	if msg.Content.Plain != nil {
		content = append(
			content,
			struct {
				Type  string
				Value string
			}{
				Type:  "text/plain",
				Value: *msg.Content.Plain,
			},
		)
	}
	if msg.Content.HTML != nil {
		content = append(
			content,
			struct {
				Type  string
				Value string
			}{
				Type:  "text/html",
				Value: *msg.Content.HTML,
			},
		)
	}
	req.Content = content

	for i := 0; i < len(msg.Attachments); i++ {
		attachments = append(
			attachments,
			Attachment{
				Content:     msg.Attachments[i].Content,
				Type:        msg.Attachments[i].Type,
				Filename:    msg.Attachments[i].Filename,
				Disposition: msg.Attachments[i].Disposition,
				ContentID:   msg.Attachments[i].ContentID,
			},
		)
	}
	req.Attachments = attachments

	return req
}
