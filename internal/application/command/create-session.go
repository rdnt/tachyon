package command

import (
	"errors"

	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/event"
)

func (s *service) CreateSession(
	id session.Id, name string, projectId project.Id,
) error {
	_, err := s.sessions.ProjectSessionByName(projectId, name)
	if err == nil {
		return errors.New("session already exists")
	} else if !errors.Is(err, ErrSessionNotFound) && err != nil {
		return err
	}

	e := event.NewSessionCreatedEvent(event.SessionCreatedEvent{
		Id:        id,
		Name:      name,
		ProjectId: projectId,
	})

	err = s.publish(e)
	if err != nil {
		return err
	}

	return nil
}
