package websocket

import (
	"fmt"

	"tachyon/internal/pkg/event"
	"tachyon/pkg/uuid"
)

func (s *Server) OnConnect(c *Conn) error {
	fmt.Println("client connected")
	uid := uuid.New()

	err := s.commands.CreateUser(uid, uid.String())
	if err != nil {
		return err
	}

	// pid := uuid.New()
	// err = s.commands.CreateProject(pid, "project-1", uid)
	// if err != nil {
	// 	return err
	// }

	c.id = uid

	return c.WriteEvent(event.ConnectedEvent{UserId: uid.String()})
}
