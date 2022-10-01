package aggregate

import (
	"fmt"

	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/event"
)

type Session struct {
	session.Session
}

func (s *Session) ProcessEvent(e event.Event) {
	switch e := e.(type) {
	case event.SessionCreatedEvent:
		s.Id = e.Id
		s.Name = e.Name
		s.ProjectId = e.ProjectId
	default:
		fmt.Println("user: unknown event", e)
	}
}
