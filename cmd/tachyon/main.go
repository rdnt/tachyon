package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-redis/redis/v9"
	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/repository/project_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/session_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/user_repository"
	"github.com/rdnt/tachyon/internal/application/query"
	"github.com/rdnt/tachyon/internal/pkg/redis/client"
	"github.com/rdnt/tachyon/internal/pkg/redis/eventbus"
	"github.com/rdnt/tachyon/internal/pkg/redis/eventstore"
	"github.com/rdnt/tachyon/pkg/uuid"
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

	uid := uuid.Nil
	err = commands.CreateUser(uid, "user-1")
	fmt.Println(err)

	pid := uuid.Nil
	err = commands.CreateProject(pid, "project-1", uid)
	fmt.Println(err)

	m := &model{
		commands:  commands,
		queries:   queries,
		projectId: pid,
	}

	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseAllMotion())

	err = p.Start()
	if err != nil {
		log.Fatal(err)
	}
}
