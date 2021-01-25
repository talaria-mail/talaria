package core

import (
	"net/mail"
)

type (
	Message struct {
		ID     int
		Flags  Flags
		Header mail.Header
		Body   []byte
	}
)
