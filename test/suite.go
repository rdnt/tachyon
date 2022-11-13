package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v9"
	"github.com/rdnt/tachyon/internal/client/application"
	"github.com/rdnt/tachyon/internal/client/remote"
	redisclient "github.com/rdnt/tachyon/internal/pkg/redis/client"
	"github.com/rdnt/tachyon/internal/pkg/redis/eventbus"
	"github.com/rdnt/tachyon/internal/pkg/redis/eventstore"
	"github.com/rdnt/tachyon/internal/server/application/command"
	"github.com/rdnt/tachyon/internal/server/application/command/repository/project_repository"
	"github.com/rdnt/tachyon/internal/server/application/command/repository/session_repository"
	"github.com/rdnt/tachyon/internal/server/application/command/repository/user_repository"
	"github.com/rdnt/tachyon/internal/server/application/query"
	"github.com/rdnt/tachyon/internal/server/websocket"
	"gotest.tools/assert"
)

type server struct {
}

func newServer(t *testing.T) *server {
	minirdb := miniredis.RunT(t)

	_, err := minirdb.XAdd("events", "1", []string{})
	assert.NilError(t, err)

	rdb := redis.NewClient(&redis.Options{
		Addr: minirdb.Addr(),
		DB:   0,
	})

	err = rdb.XDel(context.Background(), "events", "1-0").Err()
	assert.NilError(t, err)

	redisClient := redisclient.New(rdb, "events")
	eventStore := eventstore.New(redisClient)
	eventBus := eventbus.New(redisClient)

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

	return &server{}
}

type client struct {
	app *application.Application
}

func newClient(t *testing.T) *client {
	r, err := remote.New("ws://:80/ws")
	if err != nil {
		panic(err)
	}

	app, err := application.New(r)
	if err != nil {
		panic(err)
	}

	return &client{
		app: app,
	}
}
