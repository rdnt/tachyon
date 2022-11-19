package command

import (
	"errors"

	"tachyon/internal/server/application/event"
	"tachyon/pkg/uuid"
)

type CreateUserParams struct {
	Id   uuid.UUID
	Name string
}

func (s *service) CreateUser(id uuid.UUID, name string) error {
	_, err := s.users.UserByName(name)
	if err == nil {
		return errors.New("user already exists")
	} else if !errors.Is(err, ErrUserNotFound) && err != nil {
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
