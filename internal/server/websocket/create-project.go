package websocket

import (
	"tachyon/internal/pkg/event"
	"tachyon/pkg/uuid"
)

func (s *Server) CreateProject(e event.CreateProjectEvent, c *Conn) error {
	uid := c.id

	pid := uuid.New()
	pid = uuid.MustParse("wUjYcKuM7TNXQoWeWmsQYF") // FIXME: remove

	err := s.commands.CreateProject(pid, e.Name, uid)
	if err != nil {
		return err
	}

	s.Subscribe(c, pid)

	return c.WriteEvent(event.ProjectCreatedEvent{ProjectId: pid.String(), Name: e.Name})
}
