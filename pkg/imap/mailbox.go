package imap

import (
	"context"
	"errors"
	"time"

	"github.com/emersion/go-imap"
	"github.com/nsmith5/talaria/pkg/core"
)

const delimiter = `/`

type mailbox struct {
	core.Mailbox
	subscribed bool
}

func queryByID(seq imap.Seq) core.MessageQuery {
	return func(msg core.Message) bool {
		return seq.Contains(uint32(msg.ID))
	}
}
func queryByOffset(seq imap.Seq) core.MessageQuery {
	return func(msg core.Message) bool {
		return seq.Contains(uint32(msg.Offset))
	}
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
	status := imap.NewMailboxStatus(mb.Name(), items)
	status.Flags = core.AllFlags.Strings()
	status.PermanentFlags = core.AllFlags.Strings()
	{
		msg, err := mb.First(context.Background(), core.QueryUnseen)
		if err != nil {
			return nil, err
		}
		status.UnseenSeqNum = uint32(msg.Offset)
	}

	for _, name := range items {
		switch name {
		case imap.StatusMessages:
			count, err := mb.Count(context.Background(), core.QueryExists)
			if err != nil {
				return nil, err
			}
			status.Messages = uint32(count)

		case imap.StatusUidNext:
			msg, err := mb.Last(context.Background(), core.QueryExists)
			if err != nil {
				return nil, err
			}
			status.UidNext = uint32(msg.ID + 1)

		case imap.StatusUidValidity:
			status.UidValidity = 1

		case imap.StatusRecent:
			count, err := mb.Count(context.Background(), core.QueryRecent)
			if err != nil {
				return nil, err
			}
			status.Recent = uint32(count)

		case imap.StatusUnseen:
			count, err := mb.Count(context.Background(), core.QueryUnseen)
			if err != nil {
				return nil, err
			}
			status.Unseen = uint32(count)
		}
	}

	return status, nil
}

func (mb *mailbox) SetSubscribed(subscribed bool) error {
	mb.subscribed = true
	return nil
}

func (mb *mailbox) Check() error {
	return nil
}

func (mb *mailbox) ListMessages(uid bool, seqset *imap.SeqSet, items []imap.FetchItem, ch chan<- *imap.Message) error {
	defer close(ch)
	if seqset == nil {
		return errors.New(`nil sequence set`)
	}

	var query func(seq imap.Seq) core.MessageQuery
	if uid {
		query = queryByID
	} else {
		query = queryByOffset
	}

	for _, seq := range seqset.Set {
		msgs, err := mb.Find(context.Background(), query(seq))
		if err != nil {
			return err
		}

		for _, msg := range msgs {
			imsg, err := asIMAPMessage(msg, items)
			if err != nil {
				return err
			}
			ch <- imsg
		}
	}
	return nil
}

func (mb *mailbox) SearchMessages(uid bool, criteria *imap.SearchCriteria) ([]uint32, error) {
	return nil, errors.New("not implimented")
}

func (mb *mailbox) CreateMessage(flags []string, date time.Time, body imap.Literal) error {
	return errors.New("not implimented")
}

func (mb *mailbox) UpdateMessagesFlags(uid bool, seqset *imap.SeqSet, operation imap.FlagsOp, flags []string) error {
	if seqset == nil {
		return errors.New(`nil sequence set`)
	}

	var query func(seq imap.Seq) core.MessageQuery
	if uid {
		query = queryByID
	} else {
		query = queryByOffset
	}

	for _, seq := range seqset.Set {
		msgs, err := mb.Find(context.Background(), query(seq))
		if err != nil {
			return err
		}

		for _, msg := range msgs {
			imsg, err := asIMAPMessage(msg, items)
			if err != nil {
				return err
			}
			ch <- imsg
		}
	}
	return nil

	return errors.New("not implimented")
}

func (mb *mailbox) CopyMessages(uid bool, seqset *imap.SeqSet, dest string) error {
	return errors.New("not implimented")
}

func (mb *mailbox) Expunge() error {
	return errors.New("not implimented")
}
