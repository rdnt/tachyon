package websocket

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"tachyon/internal/pkg/event"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"tachyon/internal/server/application/command"
	"tachyon/internal/server/application/query"
	wsevent "tachyon/internal/server/websocket/event"
)

type Server struct {
	upgrader websocket.Upgrader
	commands command.Service
	queries  query.Service
}

func New(commands command.Service, queries query.Service) *Server {
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
	}
}

type Conn struct {
	mux  sync.Mutex
	conn *websocket.Conn
	ctx  context.Context
}

func (c *Conn) Set(k string, v string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.ctx = context.WithValue(c.ctx, k, v)
}

func (c *Conn) Get(k string) string {
	c.mux.Lock()
	defer c.mux.Unlock()

	return c.ctx.Value(k).(string)
}

func (c *Conn) Write(b []byte) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	return c.conn.WriteMessage(websocket.TextMessage, b)
}

func (c *Conn) WriteEvent(e event.Event) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	b, err := json.Marshal(e)
	if err != nil {
		return err
	}

	var tmp map[string]any
	err = json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	tmp["type"] = e.Type()

	b, err = json.Marshal(tmp)
	if err != nil {
		return err
	}

	return c.conn.WriteMessage(websocket.TextMessage, b)
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

	err = s.HandleEvent(event.ConnectedEvent{UserId: "test"}, conn)
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

		var evt wsevent.Event
		err = json.Unmarshal(b, &evt)
		if err != nil {
			fmt.Println(err)
			continue
		}

		e, err := wsevent.FromJSON(wsevent.Type(evt.Event), b)
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
	case event.ConnectedEvent:
		return s.OnConnect(conn)

	// case wsevent.CreateUserEvent:
	// 	return s.CreateUser(e, conn)
	case wsevent.CreateProjectEvent:
		return s.CreateProject(e, conn)
	case wsevent.DrawPixelEvent:
		return s.DrawPixel(e, conn)
	default:
		return errors.New("no event handler")
	}
}
