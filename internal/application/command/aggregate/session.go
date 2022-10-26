package aggregate

import (
	"fmt"

	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/event"
	"golang.org/x/exp/slices"
)

type Session struct {
	session.Session
}

func (s *Session) ProcessEvent(e event.Event) {
	switch e := e.(type) {
	case event.SessionCreatedEvent:
		s.Id = e.SessionId
		s.Name = e.Name
		s.ProjectId = e.ProjectId
		s.UserIds = e.UserIds
	case event.JoinedSessionEvent:
		s.UserIds = append(s.UserIds, e.UserId)
	case event.LeftSessionEvent:
		idx := slices.Index(s.UserIds, e.UserId)
		s.UserIds = slices.Delete(s.UserIds, idx, idx+1)
	default:
		fmt.Println("session: unknown event", e)
	}
}
