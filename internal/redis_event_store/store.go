package event_store

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/log"
	"github.com/rdnt/tachyon/pkg/broker"
	"github.com/sanity-io/litter"
)

type FanOutExchange[E any] interface {
	Publish(event E) error
	Subscribe() (chan E, error)
}

type Store struct {
	exchange FanOutExchange[[]byte]
}

type redisEvent struct {
	Event         event.EventRWIface
	Type          event.Type
	AggregateType event.AggregateType
	AggregateId   uuid.UUID
}

func (s *Store) Subscribe(h func(e event.EventIface)) (dispose func(), err error) {
	byts, err := s.exchange.Subscribe()
	if err != nil {
		return nil, err
	}

	done := make(chan bool)

	dispose = func() {
		close(byts)
		<-done
	}

	go func() {
		for byt := range byts {
			var evt redisEvent

			err := json.Unmarshal(byt, &evt)
			if err != nil {
				log.Error(err)
				continue
			}

			e := evt.Event.(event.EventRWIface)
			e.SetType(evt.Type)
			e.SetAggregateType(evt.AggregateType)
			e.SetAggregateId(evt.AggregateId)

			fmt.Println("received event", e)

			h(e)
		}

		done <- true
	}()

	return dispose, nil
}

func (s *Store) Events() ([]event.EventIface, error) {
	return s.events, nil
}

func (s *Store) Publish(e event.EventIface) error {
	s.events = append(s.events, e)
	s.broker.Publish(e)

	return nil
}

func (s *Store) String() string {
	type storedEvent struct {
		Type          event.Type
		AggregateType event.AggregateType
		AggregateId   uuid.UUID

		Event event.EventIface
	}

	var events []storedEvent

	for _, e := range s.events {
		events = append(events, storedEvent{
			Type:          e.Type(),
			AggregateType: e.AggregateType(),
			AggregateId:   e.AggregateId(),
			Event:         e,
		})
	}

	return litter.Sdump(events)
	//b, err := json.Marshal(events)
	//if err != nil {
	//	return "error"
	//}

	//return string(b)
}

func New() *Store {
	return &Store{
		broker: broker.New[event.EventIface](),
	}
}
