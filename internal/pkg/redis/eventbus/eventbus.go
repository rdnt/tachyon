package eventbus

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/pkg/interfaces"
	redisevent "github.com/rdnt/tachyon/internal/pkg/redis/event"
)

var _ interfaces.EventBus[event.Event] = (*RedisEventBus)(nil)

type RedisEventBus struct {
	client    *redis.Client
	streamKey string
}

func New(client *redis.Client, streamKey string) *RedisEventBus {
	return &RedisEventBus{
		client:    client,
		streamKey: streamKey,
	}
}

func (r *RedisEventBus) Publish(e event.Event) error {
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

func (r *RedisEventBus) Subscribe(handler func(event2 event.Event)) (func(), error) {
	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan bool)

	dispose := func() {
		cancel()
		<-done
		return
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				done <- true
				return
			default:
				streams, err := r.client.XRead(ctx, &redis.XReadArgs{
					Streams: []string{r.streamKey, "$"},
					Block:   100 * time.Millisecond,
				}).Result()
				if err != nil {
					continue
				}

				if len(streams) == 0 {
					continue
				}

				for _, msg := range streams[0].Messages {
					e, err := r.parseEvent(msg)
					if err != nil {
						fmt.Println(err)
						continue
					}

					go handler(e)
				}
			}
		}
	}()

	return dispose, nil
}

func (r *RedisEventBus) Events() ([]event.Event, error) {
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

func (r *RedisEventBus) parseEvent(msg redis.XMessage) (event.Event, error) {
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
