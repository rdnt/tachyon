package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-redis/redis/v9"
	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/repository/project_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/session_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/user_repository"
	"github.com/rdnt/tachyon/internal/application/query"
	"github.com/rdnt/tachyon/internal/pkg/redis/eventbus"
	"github.com/rdnt/tachyon/internal/pkg/redis/eventstore"
	"github.com/rdnt/tachyon/pkg/uuid"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
		DB:   0,
	})

	const redisStreamKey = "events"
	eventStore := eventstore.New(rdb, redisStreamKey)
	eventBus := eventbus.New(rdb, redisStreamKey)

	sessionRepo, err := session_repository.New(eventStore)

	userRepo, err := user_repository.New(eventStore)

	projectRepo, err := project_repository.New(eventStore)

	commands := command.New(
		eventStore,
		eventBus,
		sessionRepo,
		projectRepo,
		userRepo,
	)

	sessionView, err := session_repository.New(eventBus)

	userView, err := user_repository.New(eventBus)

	projectView, err := project_repository.New(eventBus)
	if err != nil {
		panic(err)
	}

	queries := query.New(
		eventBus,
		sessionView,
		userView,
		projectView,
	)

	uid := uuid.Nil
	err = commands.CreateUser(uid, "user-1")

	pid := uuid.Nil
	err = commands.CreateProject(pid, "project-1", uid)

	m := &model{
		commands:  commands,
		queries:   queries,
		projectId: pid,
	}

	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseAllMotion())

	err = p.Start()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
