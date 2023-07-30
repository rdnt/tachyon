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
	pid = uuid.MustParse("wUjYcKuM7TNXQoWeWmsQYF") // FIXME: remove

	err = s.commands.CreateProject(pid, e.Name, uid)
	if err != nil {
		// FIXME: HANDLE
		// return err
	}

	return c.WriteEvent(event.ProjectCreatedEvent{ProjectId: pid.String()})
}
