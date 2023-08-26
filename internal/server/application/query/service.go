package query

import (
	"tachyon/internal/server/application/domain/project"
	"tachyon/internal/server/application/domain/session"
	"tachyon/internal/server/application/domain/user"
	"tachyon/internal/server/application/event"
	"tachyon/pkg/uuid"
)

type EventBus interface {
	Subscribe(handler func(event event.Event)) (func(), error)
	Events() (events []event.Event, err error)
}

type SessionView interface {
	ProcessEvents(events ...event.Event)
	Session(id uuid.UUID) (session.Session, error)
	//CreateSession(s session.Session) error
	//UpdateSession(s session.Session) error
}

type UserView interface {
	ProcessEvents(events ...event.Event)
	User(id uuid.UUID) (user.User, error)

	//CreateUser(u user.User) error
}

type ProjectView interface {
	ProcessEvents(events ...event.Event)
	Project(id uuid.UUID) (project.Project, error)
	//CreateUser(u user.User) error
}

//type Service interface {
//	Session(id uuid.UUID) (session.Session, error)
//	//SessionUpdated(id uuid.UUID) (chan session.Session, error)
//	User(id uuid.UUID) (user.User, error)
//	Project(id uuid.UUID) (project.Project, error)
//}

type Queries struct {
	store    EventBus
	sessions SessionView
	users    UserView
	projects ProjectView
	dispose  func()
}

func (q *Queries) Session(id uuid.UUID) (session.Session, error) {
	return q.sessions.Session(id)
}

func (q *Queries) User(id uuid.UUID) (user.User, error) {
	return q.users.User(id)
}

func (q *Queries) Project(id uuid.UUID) (project.Project, error) {
	return q.projects.Project(id)
}

func New(events EventBus, sessions SessionView, users UserView, projects ProjectView) *Queries {
	s := &Queries{
		store:    events,
		sessions: sessions,
		users:    users,
		projects: projects,
	}

	//go func() {
	//	for {
	//		func() {
	//			store, err := s.store.Subscribe()
	//			if err != nil {
	//				log.Error(err)
	//				return
	//			}
	//
	//			for e := range store {
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

func (q *Queries) Start() error {
	events, err := q.store.Events()
	if err != nil {
		return err
	}

	q.processEvents(events...)

	dispose, err := q.store.Subscribe(func(e event.Event) {
		q.processEvents(e)
	})
	if err != nil {
		return err
	}

	q.dispose = dispose

	return nil
}

func (q *Queries) processEvents(events ...event.Event) {
	for _, e := range events {
		switch e.AggregateType() {
		case event.User:
			q.users.ProcessEvents(e)
		case event.Project:
			q.projects.ProcessEvents(e)
		case event.Session:
			q.sessions.ProcessEvents(e)
		}
	}
}
