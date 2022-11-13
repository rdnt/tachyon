package command

import (
	"errors"

	"github.com/rdnt/tachyon/internal/server/application/domain/project"
	"github.com/rdnt/tachyon/internal/server/application/domain/session"
	"github.com/rdnt/tachyon/internal/server/application/domain/user"
	"github.com/rdnt/tachyon/internal/server/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type EventStore interface {
	Publish(event event.Event) error
}

type EventBus interface {
	Publish(event event.Event) error
}

var ErrSessionNotFound = errors.New("session not found")

type SessionRepository interface {
	Session(id uuid.UUID) (session.Session, error)
	ProjectSessionByName(pid uuid.UUID, name string) (session.Session, error)
	ProjectSessions(pid uuid.UUID) ([]session.Session, error)
}

var ErrProjectNotFound = errors.New("project not found")

type ProjectRepository interface {
	Project(id uuid.UUID) (project.Project, error)
	UserProjectByName(userId uuid.UUID, name string) (project.Project, error)
}

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	User(id uuid.UUID) (user.User, error)
	UserByName(name string) (user.User, error)
}

type Service interface {
	CreateUser(id uuid.UUID, name string) error
	CreateProject(id uuid.UUID, name string, ownerId uuid.UUID) error
	CreateSession(id uuid.UUID, name string, projectId uuid.UUID) error
	JoinSession(id uuid.UUID, uid uuid.UUID) error
	LeaveSession(id uuid.UUID, uid uuid.UUID) error

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
	return s.store.Publish(e)
	//err := s.store.Publish(e)
	//if err != nil {
	//	return err
	//}
	//
	//return s.bus.Publish(e)
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
