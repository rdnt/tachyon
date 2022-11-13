package event_bus

import (
	"fmt"

	"github.com/rdnt/tachyon/internal/server/application/event"
	"github.com/rdnt/tachyon/pkg/fanout"
)

type Bus struct {
	exchange fanout.Exchange[event.Event]
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

func New() *Bus {
	return &Bus{
		exchange: fanout.New[event.Event](),
	}
}
