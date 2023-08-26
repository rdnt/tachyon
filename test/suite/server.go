package suite

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v9"
	"gotest.tools/assert"

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

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc
	done   chan bool
	mini   *miniredis.Miniredis
	server *http.Server
}

func NewServer(t *testing.T) *Server {
	s := &Server{
		done: make(chan bool),
	}

	s.ctx, s.cancel = context.WithCancel(context.Background())

	s.mini = miniredis.RunT(t)

	// make sure events stream is created
	//_, err := s.mini.XAdd("events", "1", []string{})
	//assert.NilError(t, err)

	rdb := redis.NewClient(&redis.Options{
		Addr: s.mini.Addr(),
		DB:   0,
	})

	err := rdb.XAdd(s.ctx, &redis.XAddArgs{
		Stream: "events",
		MaxLen: 0,
		ID:     "1-0",
		Values: map[string]interface{}{
			"event": "",
		},
	}).Err()
	assert.NilError(t, err)

	err = rdb.XDel(s.ctx, "events", "1-0").Err()
	assert.NilError(t, err)

	redisClient := redisclient.New(rdb, "events")

	eventStore := rediseventstore.New(redisClient)
	eventBus := rediseventbus.New(redisClient)

	sessionRepo, err := session_repository.New(eventStore)
	assert.NilError(t, err)

	userRepo, err := user_repository.New(eventStore)
	assert.NilError(t, err)

	projectRepo, err := project_repository.New(eventStore)
	assert.NilError(t, err)

	commands := command.NewHandler(
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

	w := websocket.New(commands, queries)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ws", w.HandlerFunc)

	s.server = &http.Server{
		Addr:    "localhost:80",
		Handler: serveMux,
	}

	go func() {
		err = s.server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			assert.NilError(t, err)
		}

		s.done <- true
	}()

	return s
}

func (s *Server) Close() error {
	s.cancel()
	_ = s.server.Close()
	<-s.done

	s.mini.Close()

	return nil
}
