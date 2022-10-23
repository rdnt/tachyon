package redis_event_bus

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/log"
)

type RedisEvent struct {
	Payload       []byte
	Type          event.Type
	AggregateType event.AggregateType
	AggregateId   uuid.UUID
}

type FanOutExchange[E any] interface {
	Publish(event E) error
	Subscribe() (chan E, error)
}

type Bus struct {
	exchange FanOutExchange[[]byte]
}

func (b *Bus) Publish(e event.EventIface) error {
	payload, err := json.Marshal(e)
	if err != nil {
		return err
	}

	evt := RedisEvent{
		Payload:       payload,
		Type:          e.Type(),
		AggregateType: e.AggregateType(),
		AggregateId:   e.AggregateId(),
	}

	byt, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	return b.exchange.Publish(byt)
}

func (b *Bus) Subscribe() (chan event.EventIface, error) {
	byts, err := b.exchange.Subscribe()
	if err != nil {
		return nil, err
	}

	evts := make(chan event.EventIface)

	go func() {
		for byt := range byts {
			var evt RedisEvent

			err := json.Unmarshal(byt, &evt)
			if err != nil {
				log.Error(err)
				continue
			}

			switch evt.Type {
			case event.UserCreated:
				var e event.UserCreatedEvent
				e.SetType(evt.Type)
				e.SetAggregateType(evt.AggregateType)
				e.SetAggregateId(evt.AggregateId)

				err := json.Unmarshal(evt.Payload, &e)
				if err != nil {
					log.Error(err)
					continue
				}

				evts <- e
			case event.ProjectCreated:
				var e event.ProjectCreatedEvent
				e.SetType(evt.Type)
				e.SetAggregateType(evt.AggregateType)
				e.SetAggregateId(evt.AggregateId)

				err := json.Unmarshal(evt.Payload, &e)
				if err != nil {
					log.Error(err)
					continue
				}

				evts <- e
			case event.PixelDrawn:
				var e event.PixelDrawnEvent
				e.SetType(evt.Type)
				e.SetAggregateType(evt.AggregateType)
				e.SetAggregateId(evt.AggregateId)

				err := json.Unmarshal(evt.Payload, &e)
				if err != nil {
					log.Error(err)
					continue
				}

				evts <- e
			default:
				fmt.Println("invalid event", evt)
			}

			//e := evt.Event.(event.EventRWIface)
			//e.SetType(evt.Type)
			//e.SetAggregateType(evt.AggregateType)
			//e.SetAggregateId(evt.AggregateId)

		}

		// dispose subscription once we're done reading
		close(byts)
	}()

	return evts, nil
}

func (b *Bus) String() string {
	return fmt.Sprint(b.exchange)
}

func New(exchange FanOutExchange[[]byte]) *Bus {
	return &Bus{
		exchange: exchange,
	}
}
