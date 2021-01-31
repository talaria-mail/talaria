package pubsub

import (
	"sync"
)

type subscriber struct {
	sync.Mutex

	handler chan interface{}
	quit    chan struct{}
	done    bool
}

func (s *subscriber) publish(evt interface{}) {
	select {
	case <-s.quit:
	case s.handler <- evt:
	default:
		// events are sent on a buffered channel. If there
		// is a slow consumer that is not processing events,
		// the buffered channel will fill and newer messages
		// are ignored.
	}
}

func (s *subscriber) close() {
	s.Lock()
	if s.done == false {
		close(s.quit)
		s.done = true
	}
	s.Unlock()
}
