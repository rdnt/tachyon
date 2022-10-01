package main

import (
	"fmt"
	"os"
	"os/signal"

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

	for i := 0; i < 48; i++ {
		go func(i int) {
			j := 0
			for {
				j++

				err = exchange.Publish([]byte(fmt.Sprintf("message-%d-%d", i, j)))
				if err != nil {
					panic(err)
				}
			}
		}(i)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
