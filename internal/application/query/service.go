package query

import (
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type EventBus interface {
	Subscribe() (chan event.Event, func(), error)
}

type SessionView interface {
	Session(id uuid.UUID) (session.Session, error)
	//CreateSession(s session.Session) error
	//UpdateSession(s session.Session) error
}

type UserView interface {
	User(id uuid.UUID) (user.User, error)
	//CreateUser(u user.User) error
}

type ProjectView interface {
	Project(id uuid.UUID) (project.Project, error)
	//CreateUser(u user.User) error
}

type Service interface {
	Session(id uuid.UUID) (session.Session, error)
	User(id uuid.UUID) (user.User, error)
	Project(id uuid.UUID) (project.Project, error)
}

type service struct {
	events   EventBus
	sessions SessionView
	users    UserView
	projects ProjectView
}

func (s *service) Session(id uuid.UUID) (session.Session, error) {
	return s.sessions.Session(id)
}

func (s *service) User(id uuid.UUID) (user.User, error) {
	return s.users.User(id)
}

func (s *service) Project(id uuid.UUID) (project.Project, error) {
	return s.projects.Project(id)
}

func New(events EventBus, sessions SessionView, users UserView, projects ProjectView) Service {
	s := &service{
		events:   events,
		sessions: sessions,
		users:    users,
		projects: projects,
	}

	//go func() {
	//	for {
	//		func() {
	//			events, err := s.events.Subscribe()
	//			if err != nil {
	//				log.Error(err)
	//				return
	//			}
	//
	//			for e := range events {
	//				log.Debug("[view] receive ", e)
	//				err = s.handleEvent(e)
	//				if err != nil {
	//					log.Error(err)
	//					continue
	//				}
	//			}
	//		}()
	//	}
	//}()

	return s
}
