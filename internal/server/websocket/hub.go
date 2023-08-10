package websocket

import (
	"fmt"

	"tachyon/internal/pkg/event"
	"tachyon/pkg/broker"
	"tachyon/pkg/uuid"
)

type Hub struct {
	conns  map[uuid.UUID]*Conn
	broker *broker.Broker[uuid.UUID, event.Event]
}

func newHub() *Hub {
	return &Hub{
		conns:  make(map[uuid.UUID]*Conn),
		broker: broker.New[uuid.UUID, event.Event](),
	}
}

func (h *Hub) AddConn(c *Conn) {
	h.conns[c.id] = c
}

func (h *Hub) Conn(id uuid.UUID) *Conn {
	return h.conns[id]
}

func (h *Hub) Subscribe(c *Conn, topic uuid.UUID) {
	fmt.Println("=== SUBSCRIBE", c.id, topic)
	dispose := h.broker.Subscribe(topic, func(e event.Event) {
		c.WriteEvent(e)
	})

	c.disposeFuncs = append(c.disposeFuncs, dispose)
}

func (h *Hub) Publish(topic uuid.UUID, e event.Event) {
	fmt.Println("=== PUBLISH", topic, e.Type())
	h.broker.Publish(topic, e)
}
