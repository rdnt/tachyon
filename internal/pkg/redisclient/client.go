package redisclient

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/pkg/redisclient/redisevent"
)

type RedisClient struct {
	client    *redis.Client
	streamKey string
}

type Options struct {
	Client    *redis.Client
	StreamKey string
}

func New(opts Options) *RedisClient {
	return &RedisClient{
		client:    opts.Client,
		streamKey: opts.StreamKey,
	}
}

//type jsonEvent struct {
//	Type          string           `json:"type"`
//	AggregateType string           `json:"aggregateType"`
//	AggregateId   string           `json:"aggregateId"`
//	Data          interfaces.Event `json:"data"`
//}

func (r *RedisClient) Publish(e event.Event) error {
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

func (r *RedisClient) Subscribe() (chan event.Event, func(), error) {
	events := make(chan event.Event)

	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan bool)

	dispose := func() {
		cancel()
		<-done
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(events)
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

					events <- e
				}
			}
		}
	}()

	return events, dispose, nil
}

func (r *RedisClient) Events() ([]event.Event, error) {
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

func (r *RedisClient) parseEvent(msg redis.XMessage) (event.Event, error) {
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
