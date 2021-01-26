package core

import (
	"net/mail"
)

type (
	MessageQuery func(Message) bool
	Message      struct {
		ID     int
		Offset int
		Flags  Flags
		Header mail.Header
		Body   []byte
	}
)

var (
	QueryExists MessageQuery = func(msg Message) bool {
		return true
	}

	QueryUnseen MessageQuery = func(msg Message) bool {
		if msg.Flags.Has(FlagSeen) {
			return false
		}
		return true
	}

	QueryRecent MessageQuery = func(msg Message) bool {
		if msg.Flags.Has(FlagRecent) {
			return true
		}
		return false
	}
)
