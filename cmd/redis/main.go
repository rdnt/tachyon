package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
		DB:   0,
	})

	fmt.Print(rdb.XRange("events", "-", "+").String())
}
