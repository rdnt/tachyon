package eventstore

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v9"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/pkg/interfaces"
	redisevent "github.com/rdnt/tachyon/internal/pkg/redis/event"
	"github.com/rdnt/tachyon/pkg/broker"
)

var _ interfaces.EventStore[event.Event] = (*RedisEventStore)(nil)

type RedisEventStore struct {
	client    *redis.Client
	broker    *broker.Broker[event.Event]
	streamKey string
}

func New(client *redis.Client, streamKey string) *RedisEventStore {
	return &RedisEventStore{
		client:    client,
		broker:    broker.New[event.Event](),
		streamKey: streamKey,
	}
}

func (r *RedisEventStore) Publish(e event.Event) error {
	err := r.publish(e)
	if err != nil {
		return err
	}

	r.broker.Publish(e)

	return nil
}

func (r *RedisEventStore) publish(e event.Event) error {
	b, err := redisevent.EventToJSON(e)
	if err != nil {
		return err
	}

	err = r.client.XAdd(context.Background(), &redis.XAddArgs{
		Stream: r.streamKey,
		MaxLen: 0,
		ID:     "*",
		Values: map[string]interface{}{
			"type": e.Type().String(),
			//"aggregateType": string(e.AggregateType()),
			//"aggregateId":   e.AggregateId().String(),
			"event": string(b),
		},
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisEventStore) Subscribe(handler func(event.Event)) (func(), error) {
	return r.broker.Subscribe(handler), nil
}

func (r *RedisEventStore) Events() ([]event.Event, error) {
	msgs, err := r.client.XRange(context.Background(), r.streamKey, "-", "+").Result()
	if err != nil {
		return nil, err
	}

	events := make([]event.Event, 0, len(msgs))
	for _, msg := range msgs {
		e, err := r.parseEvent(msg)
		if err != nil {
			fmt.Println(err)
			continue
		}

		events = append(events, e)
	}

	return events, nil
}

func (r *RedisEventStore) parseEvent(msg redis.XMessage) (event.Event, error) {
	val, ok := msg.Values["event"]
	if !ok {
		return nil, errors.New("event value does not exist")
	}

	str, ok := val.(string)
	if !ok {
		return nil, errors.New("event not a string")
	}

	v2, ok := msg.Values["type"]
	if !ok {
		return nil, errors.New("event value does not exist")
	}

	s2, ok := v2.(string)
	if !ok {
		return nil, errors.New("event not a string")
	}

	typ, err := event.TypeFromString(s2)
	if err != nil {
		return nil, err
	}

	return redisevent.EventFromJSON(typ, []byte(str))
}
