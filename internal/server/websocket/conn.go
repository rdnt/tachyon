package websocket

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
	"tachyon/internal/pkg/event"
	"tachyon/pkg/uuid"
)

type Conn struct {
	mux          sync.Mutex
	ctx          context.Context
	conn         *websocket.Conn
	id           uuid.UUID
	disposeFuncs []func()
}

func (c *Conn) Write(b []byte) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	return c.conn.WriteMessage(websocket.TextMessage, b)
}

func (c *Conn) WriteEvent(e event.Event) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	b, err := event.ToJSON(e)
	if err != nil {
		return err
	}

	fmt.Println("send", string(b))

	return c.conn.WriteMessage(websocket.TextMessage, b)
}

func (c *Conn) Close() error {
	for _, dispose := range c.disposeFuncs {
		dispose()
	}

	return c.conn.Close()
}
