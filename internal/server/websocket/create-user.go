package websocket

import (
	"github.com/rdnt/tachyon/internal/pkg/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

func (s *Server) OnConnect(c *Conn) error {
	uid := uuid.New()

	err := s.commands.CreateUser(uid, uid.String())
	if err != nil {
		return err
	}

	c.Set("userId", uid.String())

	return c.WriteEvent(event.ConnectedEvent{UserId: uid.String()})
}
