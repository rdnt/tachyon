package websocket

import (
	wsevent "github.com/rdnt/tachyon/cmd/server/websocket/event"
	"github.com/rdnt/tachyon/pkg/uuid"
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

	//b, err := wsevent.ProjectCreatedEventToJSON(wsevent.ProjectCreatedEvent{
	//	ProjectId: pid.String(),
	//	OwnerId:   uid.String(),
	//	Name:      e.Name,
	//})
	//if err != nil {
	//	return err
	//}

	return nil
}
