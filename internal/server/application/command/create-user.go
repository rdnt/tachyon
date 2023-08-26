package command

import (
	"errors"

	"tachyon/internal/server/application/command/repository/user_repository"
	"tachyon/internal/server/application/event"
	"tachyon/pkg/uuid"
)

type CreateUserParams struct {
	Id   uuid.UUID
	Name string
}

func (s *Commands) CreateUser(id uuid.UUID, name string) error {
	_, err := s.users.UserByName(name)
	if err == nil {
		return errors.New("user already exists")
	} else if !errors.Is(err, user_repository.ErrUserNotFound) && err != nil {
		return err
	}

	e := event.UserCreatedEvent{
		UserId: id,
		Name:   name,
	}

	err = s.publish(e)
	if err != nil {
		return err
	}

	return nil
}
