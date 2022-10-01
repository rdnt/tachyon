package command

import (
	"errors"

	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
)

func (s *service) CreateProject(
	id project.Id, name string, ownerId user.Id,
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

	e := event.NewProjectCreatedEvent(event.ProjectCreatedEvent{
		Id:      id,
		Name:    name,
		OwnerId: u.Id,
	})

	err = s.publish(e)
	if err != nil {
		return err
	}

	return nil
}
