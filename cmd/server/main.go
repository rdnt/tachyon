package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-redis/redis/v9"

	"tachyon/internal/pkg/redis/redisclient"
	"tachyon/internal/pkg/redis/rediseventbus"
	"tachyon/internal/pkg/redis/rediseventstore"
	"tachyon/internal/server/application/command"
	"tachyon/internal/server/application/command/repository/project_repository"
	"tachyon/internal/server/application/command/repository/session_repository"
	"tachyon/internal/server/application/command/repository/user_repository"
	"tachyon/internal/server/application/query"
	"tachyon/internal/server/websocket"
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

	err = rdb.FlushDB(context.Background()).Err()
	if err != nil {
		log.Fatal(err)
	}

	redisClient := redisclient.New(rdb, cfg.Redis.StreamKey)

	eventStore := rediseventstore.New(redisClient)
	eventBus := rediseventbus.New(redisClient)

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

	// uid := uuid.Nil
	// err = commands.CreateUser(uid, "user-1")
	// fmt.Println(err)
	//
	// pid := uuid.Nil
	// err = commands.CreateProject(pid, "project-1", uid)
	// fmt.Println(err)

	// m := &model{
	//	commands:  commands,
	//	queries:   queries,
	//	projectId: pid,
	// }
	//
	// p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseAllMotion())
	//
	// err = p.Start()
	// if err != nil {
	//	log.Fatal(err)
	// }

	s := websocket.New(commands, queries)

	http.HandleFunc("/ws", s.HandlerFunc)

	go func() {
		http.ListenAndServe(":80", nil)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGHUP)

	<-stop
}
