package command

import (
	"errors"

	"github.com/rdnt/tachyon/internal/server/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

func (s *service) CreateSession(
	id uuid.UUID, name string, projectId uuid.UUID,
) error {
	p, err := s.projects.Project(projectId)
	if err != nil {
		return err
	}

	_, err = s.sessions.ProjectSessionByName(projectId, name)
	if err == nil {
		return errors.New("session already exists")
	} else if !errors.Is(err, ErrSessionNotFound) && err != nil {
		return err
	}

	e := event.SessionCreatedEvent{
		SessionId: id,
		Name:      name,
		ProjectId: projectId,
		UserIds:   []uuid.UUID{p.OwnerId},
	}

	err = s.publish(e)
	if err != nil {
		return err
	}

	return nil
}