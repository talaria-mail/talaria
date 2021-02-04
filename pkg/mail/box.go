package mail

import (
	"code.nfsmith.ca/nsmith/talaria/pkg/kv"
	"code.nfsmith.ca/nsmith/talaria/pkg/pubsub"
)

type Box struct {
	PubSub pubsub.PubSub
	KV     kv.Store
}

func (b Box) Run() error {
	return nil
}

func (b Box) Shutdown(error) {

}
