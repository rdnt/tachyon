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
	ProjectSessions(pid project.Id) ([]session.Session, error)
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

	DrawPixel(args DrawPixelArgs) error
	ErasePixel(args ErasePixelArgs) error
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

func New(eventStore EventStore, messageBus EventBus, sessions SessionRepository, projects ProjectRepository, users UserRepository) Service {
	return &service{
		store:    eventStore,
		bus:      messageBus,
		sessions: sessions,
		projects: projects,
		users:    users,
	}
}
