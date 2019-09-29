# Technical Overview

The core logic of Talaria is to access, send and receive email. We'd like to
separate this core functionality from the protocols we'll layer over top later
on. Some time spent here designing a good interface will help us down the road.

Lets say we name the core package email. What data structures and interface do
we need here to build a robust set of functionality around?

We need a message structure:

```Go
type Message struct {
    Headers
    Body
}
```

And we'd like to flush out those MTA and MSA interfaces:

```
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

> Note: Maybe the best approach here is to simply implement the IMAP interface
> first and then try to abstract post mortem?
