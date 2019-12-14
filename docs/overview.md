# Technical Overview

The core logic of Talaria is to access, send and receive email. We'd like to
separate this core functionality from the protocols we'll layer over top later
on. Some time spent here designing a good interface will help us down the road.

Lets say we name the core package email. What data structures and interface do
we need here to build a robust set of functionality around?

We need a message structure:

```Go
import "net/mail"

type Message struct {
    // Mutable metadata
    Metadata struct {
        Mailbox string
        ID uint
        Flags uint
        // Probably more fields here...
    }

    // Immutable message
    Header mail.Header
    Body []byte
}
```

And we'd like to flush out those MTA and MSA interfaces:

```Go
type Sender interface {
    Send(m Message) error
}

type Receiver interface {
    Receive(m Message) error
}

type SendFunc func(m Message) error
type ReceiveFunc func(m Message) error
```

These interfaces will be important because we can use them to write message
middleware. For example, we can write DKIM signing outbound middleware like:

```Go
func DKIMSign(key []byte, s Sender) Sender {
    return func (m Message) error {
        m := dkim.Sign(key, m)
        return s.Send(m)
    }
}
```

Or, we can write inbound middleware like a spam detector:

```Go
func DropSpam(d SpamDetector, r Receiver) Receiver {
    return func (m Message) error {
        if d.IsSpam(m) {
            return errors.New("Spam!")
        }
        return r.Receive(m)
    }
}
```

OK, so these core interfaces are helpful. They help us abstract the core idea
behind a mail transfer agent (MTA) with `email.Sender` and they help us
abstract the idea behind a mail submission agent (MSA) with `email.Receiver`,
but there is a big gap still in functionality. We need to abstract the generic
reading and deleting interface that the IMAP protocol presents to mail clients.

> Note that I say only 'read' and 'delete' there. This is not a full CRUD
> layer, because mail (at least the message) is immutable. It is created by the
> `email.Receiver` and never updated.

This is tricky because the interface must be rich enough to implement something
as complex as the IMAP protocol.

With a problem like this, perhaps it is best to defer the trouble of complexity
to the caller.  Who knows what functionality IMAP could need? How do we ensure
that we can deal with extensions in the future? Generically, we need to select
and delete messages. Lets allow those functions to take a filter function so
they remain fairly flexibly:

```Go
type FilterFunc func(m Message) bool

type Repository interface {
    Receiver
    Select(f FilterFunc) chan Message
    Delete(f FilterFunc) uint
    Count(f FilterFunc) uint
}
```

Hmm. Ok we have an `email.Repository`. Seems like a good foundation for now.
