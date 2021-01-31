package talaria

const (
	// Inbound signals a message the outside world to a users mailbox
	Inbound int = iota

	// Outbound signals a message from a user to the outside world
	Outbound

	// Updated signals the modifcation of a message
	Updated

	// Deleted signals the deletion of a message
	Deleted
)

// Event describes a change to one message
type Event struct {
	Type    int
	Message *Message
}
