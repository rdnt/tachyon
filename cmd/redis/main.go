package main

import (
	"fmt"
	"github.com/rdnt/tachyon/internal/pkg/redisclient"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
		DB:   0,
	})

	rc := redisclient.New(redisclient.Options{
		Client:    rdb,
		StreamKey: "events",
	})

	// rc.Publish("hello :D")
	events, dispose, err := rc.Subscribe()
	if err != nil {
		panic(err)
	}
	defer dispose()

	go func() {
		for e := range events {
			fmt.Println(e)
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		err := rc.Publish("hello :D")
		if err != nil {
			panic(err)
		}

		fmt.Println(rc.Events())
		time.Sleep(1 * time.Second)

		err = rc.Publish("hello again :D")
		if err != nil {
			panic(err)
		}

		fmt.Println(rc.Events())
	}()

	<-c
}
