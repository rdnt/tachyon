package command

import (
	"errors"

	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

func (s *service) CreateProject(
	id uuid.UUID, name string, ownerId uuid.UUID,
) error {
	u, err := s.users.User(ownerId)
	if err != nil {
		return err
	}

	_, err = s.projects.UserProjectByName(u.Id, name)
	if err == nil {
		return errors.New("project already exists")
	} else if !errors.Is(err, ErrProjectNotFound) && err != nil {
		return err
	}

	e := event.ProjectCreatedEvent{
		ProjectId: id,
		Name:      name,
		OwnerId:   u.Id,
	}

	err = s.publish(e)
	if err != nil {
		return err
	}

	return nil
}
