package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/pkg/redisclient"
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
		err := rc.Publish(event.NewPixelDrawnEvent(event.PixelDrawnEvent{
			UserId:    uuid.UUID{},
			ProjectId: uuid.UUID{},
			Color:     project.Color{},
			Coords:    project.Vector2{},
		}))
		if err != nil {
			panic(err)
		}
		time.Sleep(10 * time.Millisecond)

		fmt.Println(rc.Events())
		time.Sleep(1 * time.Second)

		fmt.Println(rc.Events())
	}()

	<-c
}
