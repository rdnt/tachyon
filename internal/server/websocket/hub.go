package websocket

import (
	"tachyon/internal/pkg/event"
	"tachyon/pkg/broker"
	"tachyon/pkg/uuid"
)

type Hub struct {
	conns  map[uuid.UUID]*Conn
	broker *broker.Broker[event.Type, event.Event]
}

func newHub() *Hub {
	return &Hub{
		conns:  make(map[uuid.UUID]*Conn),
		broker: broker.New[event.Type, event.Event](),
	}
}

func (h *Hub) AddConn(c *Conn) {
	h.conns[c.id] = c
}

func (h *Hub) Conn(id uuid.UUID) *Conn {
	return h.conns[id]
}

func (h *Hub) Subscribe(id uuid.UUID, typ event.Type) {
	disp := h.broker.Subscribe(typ, func(e event.Event) {

	}
}
