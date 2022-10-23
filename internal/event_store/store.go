package event_store

import (
	"sync"

	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/pkg/broker"
	"github.com/sanity-io/litter"
)

type Store struct {
	mux    sync.Mutex
	events []event.EventIface
	broker *broker.Broker[event.EventIface]
}

func (s *Store) Subscribe(h func(e event.EventIface)) (dispose func(), err error) {
	return s.broker.Subscribe(h), nil
}

func (s *Store) Events() ([]event.EventIface, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	return s.events, nil
}

func (s *Store) Publish(e event.EventIface) error {
	s.mux.Lock()
	s.events = append(s.events, e)
	s.mux.Unlock()

	s.broker.Publish(e)

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

		Event event.EventIface
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
		events: make([]event.EventIface, 0),
		broker: broker.New[event.EventIface](),
	}
}
