package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/repository/project_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/session_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/user_repository"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/application/query"
	"github.com/rdnt/tachyon/internal/event_bus"
	"github.com/rdnt/tachyon/internal/event_store"
	"github.com/rdnt/tachyon/pkg/fanout"
)

func main() {
	eventBus := event_bus.New(fanout.New[event.Event]())

	eventStore := event_store.New()

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

	sessionView, err := session_repository.New(eventStore)

	userView, err := user_repository.New(eventStore)

	projectView, err := project_repository.New(eventStore)

	queries := query.New(
		eventBus,
		sessionView,
		userView,
		projectView,
	)

	uid := user.Id(uuid.New())
	err = commands.CreateUser(uid, "user-1")

	pid := project.Id(uuid.New())
	err = commands.CreateProject(pid, "project-1", uid)

	m := &model{
		commands:  commands,
		queries:   queries,
		userId:    uid,
		projectId: pid,
	}

	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseAllMotion())

	err = p.Start()
	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}