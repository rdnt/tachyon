package main_test

import (
	"testing"
	"time"

	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/repository/project_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/session_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/user_repository"
	"github.com/rdnt/tachyon/internal/application/query"
	"github.com/rdnt/tachyon/internal/event_store"
	"gotest.tools/assert"
)

type suite struct {
	bus         *event_store.Store
	store       *event_store.Store
	sessionRepo *session_repository.Repo
	userRepo    *user_repository.Repo
	projectRepo *project_repository.Repo
	commands    command.Service
	sessionView *session_repository.Repo
	userView    *user_repository.Repo
	projectView *project_repository.Repo
	queries     query.Service
}

func newSuite(t *testing.T) *suite {
	//eventBus := event_bus.New(fanout.New[event.Event]())

	eventStore := event_store.New()
	eventBus := eventStore

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

	sessionView, err := session_repository.New(eventBus)
	assert.NilError(t, err)

	userView, err := user_repository.New(eventBus)
	assert.NilError(t, err)

	projectView, err := project_repository.New(eventBus)
	assert.NilError(t, err)
	//sessionView := session_view.New()
	//userView := user_view.New()

	queries := query.New(
		eventBus,
		sessionView,
		userView,
		projectView,
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
		projectView: projectView,
		queries:     queries,
	}
}

func eventually(t *testing.T, f func() bool) {
	t.Helper()

	ch := make(chan bool, 1)

	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()

	ticker := time.NewTicker(1 * time.Nanosecond)
	defer ticker.Stop()

	for tick := ticker.C; ; {
		select {
		case <-timer.C:
			t.Fail()
			return
		case <-tick:
			tick = nil
			go func() { ch <- f() }()
		case v := <-ch:
			if v {
				return
			}
			tick = ticker.C
		}
	}
}
