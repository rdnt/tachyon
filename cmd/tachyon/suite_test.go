package main_test

import (
	"testing"

	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/repository/project_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/session_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/user_repository"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/application/query"
	"github.com/rdnt/tachyon/internal/application/query/view/session_view"
	"github.com/rdnt/tachyon/internal/application/query/view/user_view"
	"github.com/rdnt/tachyon/internal/event_bus"
	"github.com/rdnt/tachyon/internal/event_store"
	"github.com/rdnt/tachyon/pkg/fanout"
	"gotest.tools/assert"
)

type suite struct {
	bus         *event_bus.Bus
	store       *event_store.Store
	sessionRepo *session_repository.Repo
	userRepo    *user_repository.Repo
	projectRepo *project_repository.Repo
	commands    command.Service
	sessionView *session_view.View
	userView    *user_view.View
	queries     query.Service
}

func newSuite(t *testing.T) *suite {
	eventBus := event_bus.New(fanout.New[event.Event]())

	eventStore := event_store.New()

	sessionRepo, err := session_repository.New(eventStore)
	assert.NilError(t, err)

	userRepo, err := user_repository.New(eventStore)
	assert.NilError(t, err)

	projectRepo, err := project_repository.New(eventStore)
	assert.NilError(t, err)

	commands := command.New(
		eventStore,
		eventBus,
		sessionRepo,
		projectRepo,
		userRepo,
	)

	sessionView := session_view.New()
	userView := user_view.New()

	queries := query.New(
		eventBus,
		sessionView,
		userView,
	)

	return &suite{
		bus:         eventBus,
		store:       eventStore,
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
		projectRepo: projectRepo,
		commands:    commands,
		sessionView: sessionView,
		userView:    userView,
		queries:     queries,
	}
}
