package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-redis/redis/v9"
	"github.com/rdnt/tachyon/cmd/server/websocket"
	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/repository/project_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/session_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/user_repository"
	"github.com/rdnt/tachyon/internal/application/query"
	"github.com/rdnt/tachyon/internal/pkg/redis/client"
	"github.com/rdnt/tachyon/internal/pkg/redis/eventbus"
	"github.com/rdnt/tachyon/internal/pkg/redis/eventstore"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Database,
	})

	redisClient := client.New(rdb, cfg.Redis.StreamKey)
	eventStore := eventstore.New(redisClient)
	eventBus := eventbus.New(redisClient)

	sessionRepo, err := session_repository.New(eventStore)
	if err != nil {
		log.Fatal(err)
	}

	userRepo, err := user_repository.New(eventStore)
	if err != nil {
		log.Fatal(err)
	}

	projectRepo, err := project_repository.New(eventStore)
	if err != nil {
		log.Fatal(err)
	}

	commands := command.New(
		eventStore,
		eventBus,
		sessionRepo,
		projectRepo,
		userRepo,
	)

	sessionView, err := session_repository.New(eventBus)
	if err != nil {
		log.Fatal(err)
	}

	userView, err := user_repository.New(eventBus)
	if err != nil {
		log.Fatal(err)
	}

	projectView, err := project_repository.New(eventBus)
	if err != nil {
		log.Fatal(err)
	}

	queries := query.New(
		eventBus,
		sessionView,
		userView,
		projectView,
	)

	s := websocket.New(commands, queries)

	http.HandleFunc("/ws", s.HandlerFunc)

	go func() {
		http.ListenAndServe(":80", nil)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGHUP)

	<-stop
}
