package redis_fanout_exchange

import (
	"github.com/go-redis/redis"
)

type Exchange struct {
	client *redis.Client
}

func (e *Exchange) Publish(event []byte) error {
	err := e.client.XAdd(&redis.XAddArgs{
		Stream:       "events",
		MaxLen:       0,
		MaxLenApprox: 0,
		ID:           "*",
		Values: map[string]interface{}{
			"event": event,
		},
	}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (e *Exchange) Subscribe() (chan []byte, error) {
	//err := e.client.XGroupCreate("events", "query", "$").Err()
	//if err != nil && err.Error() != errBusyGroup {
	//	return nil, err
	//}

	events := make(chan []byte)

	go func() {
		for {
			func() {
				streams, err := e.client.XRead(&redis.XReadArgs{
					Streams: []string{"events", "$"},
					Count:   0,
					Block:   0,
				}).Result()

				if err != nil {
					return
				}

				if len(streams) == 0 {
					return
				}

				for _, msg := range streams[0].Messages {
					evt, ok := msg.Values["event"]
					if !ok {
						continue
					}

					str, ok := evt.(string)
					if !ok {
						continue
					}

					events <- []byte(str)
				}
			}()
		}
	}()

	return events, nil
}

func New(client *redis.Client) *Exchange {
	return &Exchange{
		client: client,
	}
}
