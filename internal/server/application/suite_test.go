package application_test

import (
	"testing"
	"time"

	"gotest.tools/assert"

	"tachyon/internal/server/application/command"
	"tachyon/internal/server/application/command/repository/project_repository"
	"tachyon/internal/server/application/command/repository/session_repository"
	"tachyon/internal/server/application/command/repository/user_repository"
	"tachyon/internal/server/application/event"
	"tachyon/internal/server/application/event_store"
	"tachyon/internal/server/application/query"
)

type EventStore interface {
	Publish(event event.Event) error
	Subscribe(h func(e event.Event)) (dispose func(), err error)
	Events() (events []event.Event, err error)
}

type suite struct {
	bus         *event_store.Store
	store       *event_store.Store
	sessionRepo *session_repository.Repo
	userRepo    *user_repository.Repo
	projectRepo *project_repository.Repo
	commands    *command.Commands
	sessionView *session_repository.Repo
	userView    *user_repository.Repo
	projectView *project_repository.Repo
	queries     *query.Queries
}

func newSuite(t *testing.T) *suite {
	eventStore := event_store.New()
	eventBus := eventStore

	sessionRepo, err := session_repository.New(eventStore)
	assert.NilError(t, err)

	userRepo, err := user_repository.New(eventStore)
	assert.NilError(t, err)

	projectRepo, err := project_repository.New(eventStore)
	assert.NilError(t, err)

	commandHandler := command.NewHandler(
		eventStore,
		eventBus,
		sessionRepo,
		projectRepo,
		userRepo,
	)

	err = commandHandler.Start()
	assert.NilError(t, err)

	sessionView, err := session_repository.New(eventBus)
	assert.NilError(t, err)

	userView, err := user_repository.New(eventBus)
	assert.NilError(t, err)

	projectView, err := project_repository.New(eventBus)
	assert.NilError(t, err)
	//sessionView := session_view.NewHandler()
	//userView := user_view.NewHandler()

	queries := query.New(
		eventBus,
		sessionView,
		userView,
		projectView,
	)

	err = queries.Start()
	assert.NilError(t, err)

	return &suite{
		bus:         eventBus,
		store:       eventStore,
		sessionRepo: sessionRepo,
		userRepo:    userRepo,
		projectRepo: projectRepo,
		commands:    commandHandler,
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
