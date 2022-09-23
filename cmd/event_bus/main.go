package main

import (
	"github.com/go-redis/redis"
	"github.com/rdnt/tachyon/pkg/redis_fanout_exchange"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Network: "",
		Addr:    "localhost:6379",
	})

	exchange := redis_fanout_exchange.New(rdb)

	events, err := exchange.Subscribe()
	if err != nil {
		panic(err)
	}

	go func() {
		for e := range events {
			if len(e) == 0 {
				continue
			}
		}
	}()

	for {
		err = exchange.Publish([]byte("message"))
		if err != nil {
			panic(err)
		}
	}
}
