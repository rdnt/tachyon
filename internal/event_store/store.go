package event_store

import (
	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/pkg/broker"
	"github.com/sanity-io/litter"
)

type Store struct {
	events []event.Event
	broker *broker.Broker[event.Event]
}

func (s *Store) Subscribe(h func(e event.Event)) (dispose func(), err error) {
	return s.broker.Subscribe(h), nil
}

func (s *Store) Events() ([]event.Event, error) {
	return s.events, nil
}

func (s *Store) Publish(e event.Event) error {
	s.events = append(s.events, e)
	s.broker.Publish(e)

	return nil
}

func (s *Store) String() string {
	type storedEvent struct {
		Type          event.Type
		AggregateType event.AggregateType
		AggregateId   uuid.UUID

		Event event.Event
	}

	events := []storedEvent{}

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
		events: make([]event.Event, 0),
		broker: broker.New[event.Event](),
	}
}
