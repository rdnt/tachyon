package command

import (
	"errors"

	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
)

type CreateUserParams struct {
	Id   user.Id
	Name string
}

func (s *service) CreateUser(id user.Id, name string) error {
	_, err := s.users.UserByName(name)
	if err == nil {
		return errors.New("user already exists")
	} else if !errors.Is(err, ErrUserNotFound) && err != nil {
		return err
	}

	e := event.NewUserCreatedEvent(event.UserCreatedEvent{
		Id:   id,
		Name: name,
	})

	err = s.publish(e)
	if err != nil {
		return err
	}

	return nil
}
