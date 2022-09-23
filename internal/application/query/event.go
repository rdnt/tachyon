package query

import (
	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
)

func (s *service) handleEvent(e event.Event) error {
	switch e := e.(type) {
	case event.SessionCreatedEvent:
		err := s.sessions.CreateSession(session.Session{
			Id:        e.Id,
			Name:      e.Name,
			ProjectId: e.ProjectId,
			OwnerId:   e.OwnerId,
			UserIds:   e.UserIds,
		})
		if err != nil {
			return err
		}
	case event.UserCreatedEvent:
		err := s.users.CreateUser(user.User{
			Id:   e.Id,
			Name: e.Name,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
