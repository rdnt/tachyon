package event_bus

import (
	"fmt"

	"github.com/rdnt/tachyon/internal/application/event"
)

type FanOutExchange[E any] interface {
	Publish(event E) error
	Subscribe() (chan E, error)
}

type Bus struct {
	exchange FanOutExchange[event.EventIface]
}

func (b *Bus) Publish(event event.EventIface) error {
	return b.exchange.Publish(event)
}

func (b *Bus) Subscribe() (chan event.EventIface, error) {
	return b.exchange.Subscribe()
}

func (b *Bus) String() string {
	return fmt.Sprint(b.exchange)
}

func New(exchange FanOutExchange[event.EventIface]) *Bus {
	return &Bus{
		exchange: exchange,
	}
}
