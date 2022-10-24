package redisclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/rdnt/tachyon/internal/pkg/interfaces"
	"time"
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

func (r *RedisClient) Publish(e interfaces.Event) error {
	b, err := json.Marshal(e)
	if err != nil {
		return err
	}

	err = r.client.XAdd(&redis.XAddArgs{
		Stream:       r.streamKey,
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "*",
		Values: map[string]interface{}{
			"event": string(b),
		},
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) Subscribe() (chan interfaces.Event, func(), error) {
	events := make(chan interfaces.Event)

	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan bool)

	dispose := func() {
		fmt.Println("disposing...")
		cancel()
		close(events)
		fmt.Println("disposed")
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("default stmt")
				close(events)
				done <- true
				return
			default:
				fmt.Println("listen")
				streams, err := r.client.WithContext(ctx).XRead(&redis.XReadArgs{
					Streams: []string{r.streamKey, "$"},
					Count:   0,
					Block:   10 * time.Second,
				}).Result()
				if err != nil {
					continue
				}

				fmt.Println(streams)

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

func (r *RedisClient) Events() ([]interfaces.Event, error) {
	msgs, err := r.client.XRange(r.streamKey, "-", "+").Result()
	if err != nil {
		return nil, err
	}

	events := make([]interfaces.Event, 0, len(msgs))
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

func (r *RedisClient) parseEvent(msg redis.XMessage) (interfaces.Event, error) {
	evt, ok := msg.Values["event"]
	if !ok {
		return nil, errors.New("event value does not exist")
	}

	str, ok := evt.(string)
	if !ok {
		return nil, errors.New("event not a string")
	}

	var e interfaces.Event
	err := json.Unmarshal([]byte(str), &e)
	if err != nil {
		return nil, err
	}

	return e, nil
}
