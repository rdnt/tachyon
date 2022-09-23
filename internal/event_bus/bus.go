package event_bus

import (
	"fmt"

	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/log"
)

type FanoutExhange[E any] interface {
	Publish(event E) error
	Subscribe() (chan E, error)
}

type Bus struct {
	exchange FanoutExhange[event.Event]
}

func (b *Bus) Publish(event event.Event) error {
	log.Debug("[cmds] send ", event)
	return b.exchange.Publish(event)
}

func (b *Bus) Subscribe() (chan event.Event, error) {
	return b.exchange.Subscribe()
}

func (b *Bus) String() string {
	return fmt.Sprint(b.exchange)
}

func New(exchange FanoutExhange[event.Event]) *Bus {
	return &Bus{
		exchange: exchange,
	}
}
