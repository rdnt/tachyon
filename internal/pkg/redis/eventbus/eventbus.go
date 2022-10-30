package eventbus

import (
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/pkg/redis/client"
)

type EventBus struct {
	client *client.Client
}

func New(client *client.Client) *EventBus {
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
