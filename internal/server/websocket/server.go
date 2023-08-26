package websocket

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"tachyon/pkg/uuid"

	"tachyon/internal/pkg/event"
	"tachyon/internal/server/application/command"
	"tachyon/internal/server/application/query"
)

type Server struct {
	upgrader websocket.Upgrader
	commands *command.Commands
	queries  *query.Queries
	hub      *Hub
}

func New(commands *command.Commands, queries *query.Queries) *Server {
	return &Server{
		commands: commands,
		queries:  queries,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  0,
			WriteBufferSize: 0,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			EnableCompression: false,
		},
		hub: newHub(),
	}
}

func (s *Server) HandlerFunc(w http.ResponseWriter, req *http.Request) {
	wsconn, err := s.upgrader.Upgrade(w, req, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	wsconn.EnableWriteCompression(false)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn := &Conn{
		conn: wsconn,
		ctx:  ctx,
	}

	err = s.OnConnect(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		typ, b, err := wsconn.ReadMessage()
		if err != nil {
			return
		}

		if typ != websocket.TextMessage {
			fmt.Println("invalid message")
			continue
		}

		e, err := event.FromJSON(b)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = s.HandleEvent(e, conn)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func (s *Server) HandleEvent(e any, conn *Conn) error {
	switch e := e.(type) {
	// case event.CreateUserEvent:
	// 	return s.CreateUser(e, conn)
	case event.CreateProjectEvent:
		return s.CreateProject(e, conn)
	case event.CreateSessionEvent:
		return s.CreateSession(e, conn)
	case event.CreatePathEvent:
		return s.CreatePath(e, conn)
	default:
		return errors.New("no event handler")
	}
}

func (s *Server) Publish(topic uuid.UUID, e event.Event) error {
	s.hub.Publish(topic, e)
	return nil
}

func (s *Server) Subscribe(c *Conn, topic uuid.UUID) {
	s.hub.Subscribe(c, topic)
}
