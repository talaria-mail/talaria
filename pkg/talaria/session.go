package talaria

import (
	"time"
)

// Session store context and data for an authenticated users
type Session struct {
	User       User
	Expiration time.Time

	// Decrypted content encryption key
	ContentKey []byte
}
