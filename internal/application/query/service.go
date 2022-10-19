package query

import (
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
)

type EventBus interface {
	Subscribe() (chan event.Event, error)
}

type SessionView interface {
	Session(id session.Id) (session.Session, error)
	//CreateSession(s session.Session) error
	//UpdateSession(s session.Session) error
}

type UserView interface {
	User(id user.Id) (user.User, error)
	//CreateUser(u user.User) error
}

type ProjectView interface {
	Project(id project.Id) (project.Project, error)
	//CreateUser(u user.User) error
}

type Service interface {
	Session(id session.Id) (session.Session, error)
	User(id user.Id) (user.User, error)
	Project(id project.Id) (project.Project, error)
}

type service struct {
	events   EventBus
	sessions SessionView
	users    UserView
	projects ProjectView
}

func (s *service) Session(id session.Id) (session.Session, error) {
	return s.sessions.Session(id)
}

func (s *service) User(id user.Id) (user.User, error) {
	return s.users.User(id)
}

func (s *service) Project(id project.Id) (project.Project, error) {
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
