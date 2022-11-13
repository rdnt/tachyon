package websocket

import (
	wsevent "github.com/rdnt/tachyon/cmd/server/websocket/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

func (s *Server) CreateUser(e wsevent.CreateUserEvent, c *Conn) error {
	uid := uuid.New()

	err := s.commands.CreateUser(uid, e.Name)
	if err != nil {
		return err
	}

	c.Set("userId", uid.String())

	//b, err := wsevent.UserCreatedEventToJSON(event.UserCreatedEvent{
	//	UserId: uid,
	//	Name:   e.Name,
	//})
	//if err != nil {
	//	return err
	//}

	return nil
}
