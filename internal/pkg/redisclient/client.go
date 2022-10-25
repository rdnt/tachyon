package redisclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
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

func (r *RedisClient) Publish(e []byte) error {
	b, err := json.Marshal(e)
	if err != nil {
		return err
	}

	err = r.client.XAdd(context.Background(), &redis.XAddArgs{
		Stream: r.streamKey,
		MaxLen: 0,
		ID:     "*",
		Values: map[string]interface{}{
			"event": string(b),
		},
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *RedisClient) Subscribe() (chan []byte, func(), error) {
	events := make(chan []byte)

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

func (r *RedisClient) Events() ([][]byte, error) {
	msgs, err := r.client.XRange(context.Background(), r.streamKey, "-", "+").Result()
	if err != nil {
		return nil, err
	}

	events := make([][]byte, 0, len(msgs))
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

func (r *RedisClient) parseEvent(msg redis.XMessage) ([]byte, error) {
	evt, ok := msg.Values["event"]
	if !ok {
		return nil, errors.New("event value does not exist")
	}

	str, ok := evt.(string)
	if !ok {
		return nil, errors.New("event not a string")
	}

	return []byte(str), nil
}
