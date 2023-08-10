package rediseventstore

import (
	"tachyon/internal/pkg/redis/redisclient"
	"tachyon/internal/server/application/event"
	"tachyon/pkg/broker"
)

type EventStore struct {
	client *redisclient.Client
	broker *broker.Broker[string, event.Event]
}

func New(client *redisclient.Client) *EventStore {
	return &EventStore{
		client: client,
		broker: broker.New[string, event.Event](),
	}
}

func (store *EventStore) Publish(e event.Event) error {
	err := store.client.Publish(e)
	if err != nil {
		return err
	}

	// Publish the event through the in-memory broker so that subscribers will
	// also receive the event synchronously before publish returns
	store.broker.Publish("", e)

	return nil
}

func (store *EventStore) Subscribe(h func(e event.Event)) (func(), error) {
	// Do not use the redis tcplinepoc but subscribe directly to the broker
	return store.broker.Subscribe("", h), nil
}

func (store *EventStore) Events() ([]event.Event, error) {
	return store.client.Events()
}
