package aggregate

import (
	"fmt"

	"tachyon/internal/server/application/domain/session"
	"tachyon/internal/server/application/event"
	"tachyon/pkg/broker"

	"golang.org/x/exp/slices"
)

type Session struct {
	session.Session
	broker *broker.Simple[session.Session]
}

func NewSession() *Session {
	return &Session{
		broker: broker.NewSimple[session.Session](),
	}
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
		fmt.Println("session: unknown event 3", e.Type(), e)
		return
	}

	s.broker.Publish(s.Session)
}
