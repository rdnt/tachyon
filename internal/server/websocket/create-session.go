package websocket

import (
	"tachyon/internal/pkg/event"
	"tachyon/pkg/uuid"
)

func (s *Server) CreateSession(e event.CreateSessionEvent, c *Conn) error {
	//uid := c.id

	pid := uuid.New()
	pid = uuid.MustParse("wUjYcKuM7TNXQoWeWmsQYF") // FIXME: remove

	sid := uuid.New()
	sid = uuid.MustParse("wUjYcKuM7TNXQoWeWmsQYE") // FIXME: remove
	err := s.commands.CreateSession(sid, e.Name, pid)
	if err != nil {
		// FIXME: HANDLE
		// return err
	}

	return c.WriteEvent(event.SessionCreatedEvent{
		ProjectId: pid.String(),
		SessionId: sid.String(),
		Name:      e.Name,
	})
}
