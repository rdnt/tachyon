package eventstore

import (
	"github.com/rdnt/tachyon/internal/server/application/event"
	"github.com/rdnt/tachyon/internal/pkg/redis/client"
	"github.com/rdnt/tachyon/pkg/broker"
)

type EventStore struct {
	client *client.Client
	broker *broker.Broker[event.Event]
}

func New(client *client.Client) *EventStore {
	return &EventStore{
		client: client,
		broker: broker.New[event.Event](),
	}
}

func (store *EventStore) Publish(e event.Event) error {
	err := store.client.Publish(e)
	if err != nil {
		return err
	}

	// Publish the event through the in-memory broker so that subscribers will
	// also receive the event synchronously before publish returns
	store.broker.Publish(e)

	return nil
}

func (store *EventStore) Subscribe(h func(e event.Event)) (func(), error) {
	// Do not use the redis tcplinepoc but subscribe directly to the broker
	return store.broker.Subscribe(h), nil
}

func (store *EventStore) Events() ([]event.Event, error) {
	return store.client.Events()
}
