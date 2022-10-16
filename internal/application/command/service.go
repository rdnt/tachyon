package command

import (
	"errors"

	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
)

type EventStore interface {
	Publish(event event.Event) error
}

type EventBus interface {
	Publish(event event.Event) error
}

var ErrSessionNotFound = errors.New("session not found")

type SessionRepository interface {
	Session(id session.Id) (session.Session, error)
	ProjectSessionByName(pid project.Id, name string) (session.Session, error)
}

var ErrProjectNotFound = errors.New("project not found")

type ProjectRepository interface {
	Project(id project.Id) (project.Project, error)
	UserProjectByName(userId user.Id, name string) (project.Project, error)
}

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	User(id user.Id) (user.User, error)
	UserByName(name string) (user.User, error)
}

type Service interface {
	CreateUser(id user.Id, name string) error
	CreateProject(id project.Id, name string, ownerId user.Id) error
	CreateSession(id session.Id, name string, projectId project.Id) error
	JoinSession(id session.Id, uid user.Id) error
	LeaveSession(id session.Id, uid user.Id) error

	CreatePath(args CreatePathArgs) error
}

type service struct {
	sessions SessionRepository
	projects ProjectRepository
	users    UserRepository
	store    EventStore
	bus      EventBus
}

func (s *service) publish(e event.Event) error {
	err := s.store.Publish(e)
	if err != nil {
		return err
	}

	return s.bus.Publish(e)
}

//func (s *service) CreateSession(userId uuid.UUID, projectId project.Id, sessionName string) error {
//	u, err := s.users.User(userId)
//	if err != nil {
//		return session.Session{}, err
//	}
//
//	e := event.NewSessionCreatedEvent(event.SessionCreatedEvent{
//		Id:        session.Id(uuid.New()),
//		Name:      sessionName,
//		ProjectId: projectId,
//		OwnerId:   userId,
//		UserIds:   []user.Id{userId},
//	})
//
//	err := s.store.Publish(e)
//	if err != nil {
//		return err
//	}
//
//	err = s.bus.Publish(e)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *service) CreateProject(userId user.Id, name string) error {
//	s.projects.Project()
//}
//
//func (s *service) CreateUser(name string) error {
//	e := event.NewUserCreatedEvent(event.UserCreatedEvent{
//		//Id:   user.Id(uuid.New()),
//		Id:   user.Id(name), // TODO: use uuid
//		Name: name,
//	})
//
//	err := s.store.Publish(e)
//	if err != nil {
//		return err
//	}
//
//	err = s.bus.Publish(e)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (s *service) JoinSession(userId user.Id, sessionId session.Id) error {
//	//TODO implement me
//	panic("implement me")
//}

func New(eventStore EventStore, messageBus EventBus, sessions SessionRepository, projects ProjectRepository, users UserRepository) Service {
	return &service{
		store:    eventStore,
		bus:      messageBus,
		sessions: sessions,
		projects: projects,
		users:    users,
	}
}
