package event_store

import (
	"sync"

	"tachyon/internal/server/application/event"
	"tachyon/pkg/broker"
	"tachyon/pkg/uuid"
	"github.com/sanity-io/litter"
)

type Store struct {
	mux    sync.Mutex
	events []event.Event
	broker *broker.Broker[event.Event]
}

func (s *Store) Subscribe(handler func(e event.Event)) (dispose func(), err error) {
	return s.broker.Subscribe(func(e event.Event) {
		handler(e)
	}), nil
}

func (s *Store) Events() ([]event.Event, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	return s.events, nil
}

func (s *Store) Publish(e event.Event) error {
	s.mux.Lock()
	s.events = append(s.events, e)
	s.broker.Publish(e)

	s.mux.Unlock()

	return nil
}

func (s *Store) String() string {
	s.mux.Lock()
	storedEvents := s.events
	s.mux.Unlock()

	type storedEvent struct {
		Type          event.Type
		AggregateType event.AggregateType
		AggregateId   uuid.UUID

		Event event.Event
	}

	var events []storedEvent

	for _, e := range storedEvents {
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
