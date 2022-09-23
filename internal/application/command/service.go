package command

import (
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type EventStore interface {
	Events() ([]event.Event, error)
	Publish(e event.Event) error
}

type EventBus interface {
	Publish(event event.Event) error
}

type SessionRepository interface {
	Session(id session.Id) (session.Session, error)
}

type UserRepository interface {
	User(id user.Id) (user.User, error)
}

type Service interface {
	CreateUser(name string) error
	CreateSession(userId user.Id, projectId project.Id, name string) error
}

type service struct {
	sessions SessionRepository
	users    UserRepository
	store    EventStore
	bus      EventBus
}

func (s *service) CreateSession(userId user.Id, projectId project.Id, sessionName string) error {
	e := event.NewSessionCreatedEvent(event.SessionCreatedEvent{
		Id:        session.Id(uuid.New()),
		Name:      sessionName,
		ProjectId: projectId,
		OwnerId:   userId,
		UserIds:   []user.Id{userId},
	})

	err := s.store.Publish(e)
	if err != nil {
		return err
	}

	err = s.bus.Publish(e)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CreateUser(name string) error {
	e := event.NewUserCreatedEvent(event.UserCreatedEvent{
		//Id:   user.Id(uuid.New()),
		Id:   user.Id(name), // TODO: use uuid
		Name: name,
	})

	err := s.store.Publish(e)
	if err != nil {
		return err
	}

	err = s.bus.Publish(e)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) JoinSession(userId user.Id, sessionId session.Id) error {
	//TODO implement me
	panic("implement me")
}

func New(eventStore EventStore, messageBus EventBus, sessions SessionRepository, users UserRepository) Service {
	return &service{
		store:    eventStore,
		bus:      messageBus,
		sessions: sessions,
		users:    users,
	}
}
