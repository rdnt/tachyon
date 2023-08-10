package aggregate

import (
	"fmt"

	"tachyon/internal/client/application/domain/session"
	"tachyon/internal/pkg/event"
	"tachyon/pkg/uuid"
)

type Session struct {
	session.Session
}

func (s *Session) ProcessEvent(e event.Event) {
	switch e := e.(type) {
	case event.SessionCreatedEvent:
		s.Id = uuid.MustParse(e.SessionId)
		s.Name = e.Name
		s.ProjectId = uuid.MustParse(e.ProjectId)
		//s.UserIds = e.UserIds // TODO: ??
	//case event.JoinedSessionEvent:
	//	s.UserIds = append(s.UserIds, e.UserId)
	//case event.LeftSessionEvent:
	//	idx := slices.Index(s.UserIds, e.UserId)
	//	s.UserIds = slices.Delete(s.UserIds, idx, idx+1)
	default:
		fmt.Println("session: unknown event", e)
	}
}
