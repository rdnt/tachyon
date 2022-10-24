package event_store

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/log"
	"github.com/sanity-io/litter"
)

type FanOutExchange[E any] interface {
	Publish(event E) error
	Subscribe() (chan E, error)
}

type Store struct {
	client *redis.Client
}

type redisEvent struct {
	Event         event.EventRWIface
	Type          event.Type
	AggregateType event.AggregateType
	AggregateId   uuid.UUID
}

//func (s *Store) Subscribe(h func(e event.EventIface)) (dispose func(), err error) {
//	byts, err := s.exchange.Subscribe()
//	if err != nil {
//		return nil, err
//	}
//
//	done := make(chan bool)
//
//	dispose = func() {
//		close(byts)
//		<-done
//	}
//
//	go func() {
//		for byt := range byts {
//			var evt redisEvent
//
//			err := json.Unmarshal(byt, &evt)
//			if err != nil {
//				log.Error(err)
//				continue
//			}
//
//			e := evt.Event.(event.EventRWIface)
//			e.SetType(evt.Type)
//			e.SetAggregateType(evt.AggregateType)
//			e.SetAggregateId(evt.AggregateId)
//
//			fmt.Println("received event", e)
//
//			h(e)
//		}
//
//		done <- true
//	}()
//
//	return dispose, nil
//}

func (s *Store) Events() ([]event.EventIface, error) {
	msgs, err := s.client.XRange("events", "-", "+").Result()
	if err != nil {
		return nil, err
	}

	events := make([]event.EventIface, 0, len(msgs))
	for _, msg := range msgs {
		val, ok := msg.Values["event"]
		if !ok {
			continue
		}

		str, ok := val.(string)
		if !ok {
			continue
		}

		var evt redisEvent

		err := json.Unmarshal([]byte(str), &evt)
		if err != nil {
			log.Error(err)
			continue
		}

		e := evt.Event.(event.EventRWIface)
		e.SetType(evt.Type)
		e.SetAggregateType(evt.AggregateType)
		e.SetAggregateId(evt.AggregateId)

		fmt.Println("received event", e)

		events = append(events, e)
	}

	return events, nil
}

type RedisEvent struct {
	Payload       []byte
	Type          event.Type
	AggregateType event.AggregateType
	AggregateId   uuid.UUID
}

func (s *Store) Publish(e event.EventIface) error {
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

func New(client *redis.Client) *Store {
	return &Store{
		client: client,
	}
}
