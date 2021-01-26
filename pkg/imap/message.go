package imap

import (
	"bytes"
	"net/mail"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/backend/backendutil"
	"github.com/emersion/go-message/textproto"
	"github.com/nsmith5/talaria/pkg/core"
)

func toTextProto(header mail.Header) textproto.Header {
	var output textproto.Header
	for key, values := range header {
		for _, value := range values {
			output.Add(key, value)
		}
	}
	return output
}

func asIMAPMessage(msg core.Message, items []imap.FetchItem) (*imap.Message, error) {
	var err error

	imsg := imap.NewMessage(uint32(msg.Offset), items)
	for _, item := range items {
		switch item {
		case imap.FetchEnvelope:
			imsg.Envelope, err = backendutil.FetchEnvelope(toTextProto(msg.Header))
			if err != nil {
				return nil, err
			}

		case imap.FetchBody, imap.FetchBodyStructure:
			header := toTextProto(msg.Header)
			body := bytes.NewReader(msg.Body)
			imsg.BodyStructure, err = backendutil.FetchBodyStructure(header, body, item == imap.FetchBodyStructure)
			if err != nil {
				return nil, err
			}

		case imap.FetchFlags:
			imsg.Flags = msg.Flags.Strings()

		case imap.FetchInternalDate:
			imsg.InternalDate, err = msg.Header.Date()
			if err != nil {
				return nil, err
			}

		case imap.FetchRFC822Size:
			imsg.Size = uint32(len(msg.Body))

		case imap.FetchUid:
			imsg.Uid = uint32(msg.ID)

		default:
			section, err := imap.ParseBodySectionName(item)
			if err != nil {
				break
			}

			header := toTextProto(msg.Header)
			body := bytes.NewReader(msg.Body)

			l, err := backendutil.FetchBodySection(header, body, section)
			if err != nil {
				return nil, err
			}
			imsg.Body[section] = l
		}
	}

	return imsg, nil
}
