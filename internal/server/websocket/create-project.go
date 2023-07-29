package websocket

import (
	"tachyon/internal/pkg/event"
	wsevent "tachyon/internal/server/websocket/event"
	"tachyon/pkg/uuid"
)

func (s *Server) CreateProject(e wsevent.CreateProjectEvent, c *Conn) error {
	uid, err := uuid.Parse(c.Get("userId"))
	if err != nil {
		return err
	}

	pid := uuid.New()

	err = s.commands.CreateProject(pid, e.Name, uid)
	if err != nil {
		return err
	}

	return c.WriteEvent(event.ProjectCreatedEvent{ProjectId: pid.String()})
}
