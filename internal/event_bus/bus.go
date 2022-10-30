package event_bus

import (
	"fmt"

	"github.com/rdnt/tachyon/internal/application/event"
)

type FanOutExchange[E any] interface {
	Publish(event E) error
	Subscribe() (chan E, func(), error)
}

type Bus struct {
	exchange FanOutExchange[event.Event]
}

func (b *Bus) Publish(event event.Event) error {
	return b.exchange.Publish(event)
}

func (b *Bus) Subscribe() (chan event.Event, func(), error) {
	return b.exchange.Subscribe()
}

func (b *Bus) String() string {
	return fmt.Sprint(b.exchange)
}

func New(exchange FanOutExchange[event.Event]) *Bus {
	return &Bus{
		exchange: exchange,
	}
}
