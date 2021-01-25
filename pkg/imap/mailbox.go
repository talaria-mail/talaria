package imap

import (
	"errors"
	"time"

	"github.com/emersion/go-imap"
	"github.com/nsmith5/talaria/pkg/core"
)

const delimiter = `/`

type mailbox struct {
	core.Mailbox
}

func (mb *mailbox) Info() (*imap.MailboxInfo, error) {
	info := &imap.MailboxInfo{
		Name:       mb.Mailbox.Name(),
		Delimiter:  delimiter,
		Attributes: nil,
	}
	return info, nil
}

func (mb *mailbox) Name() string {
	return mb.Mailbox.Name()
}

func (mb *mailbox) Status(items []imap.StatusItem) (*imap.MailboxStatus, error) {
	return nil, errors.New("not implimented")
}

func (mb *mailbox) SetSubscribed(subscribed bool) error {
	return errors.New("not implimented")
}

func (mb *mailbox) Check() error {
	return errors.New("not implimented")
}

func (mb *mailbox) ListMessages(uid bool, seqset *imap.SeqSet, items []imap.FetchItem, ch chan<- *imap.Message) error {
	return errors.New("not implimented")
}

func (mb *mailbox) SearchMessages(uid bool, criteria *imap.SearchCriteria) ([]uint32, error) {
	return nil, errors.New("not implimented")
}

func (mb *mailbox) CreateMessage(flags []string, date time.Time, body imap.Literal) error {
	return errors.New("not implimented")
}

func (mb *mailbox) UpdateMessagesFlags(uid bool, seqset *imap.SeqSet, operation imap.FlagsOp, flags []string) error {
	return errors.New("not implimented")
}

func (mb *mailbox) CopyMessages(uid bool, seqset *imap.SeqSet, dest string) error {
	return errors.New("not implimented")
}

func (mb *mailbox) Expunge() error {
	return errors.New("not implimented")
}
