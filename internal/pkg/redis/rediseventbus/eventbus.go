package rediseventbus

import (
	"tachyon/internal/pkg/redis/redisclient"
	"tachyon/internal/server/application/event"
)

type EventBus struct {
	client *redisclient.Client
}

func New(client *redisclient.Client) *EventBus {
	return &EventBus{client: client}
}

func (bus *EventBus) Publish(e event.Event) error {
	return bus.client.Publish(e)
}

func (bus *EventBus) Subscribe(h func(e event.Event)) (func(), error) {
	return bus.client.Subscribe(h)
}

func (bus *EventBus) Events() ([]event.Event, error) {
	return bus.client.Events()
}
