package command

import (
	"tachyon/internal/server/application/command/repository/project_repository"
	"tachyon/internal/server/application/command/repository/session_repository"
	"tachyon/internal/server/application/command/repository/user_repository"
	"tachyon/internal/server/application/event"
)

type EventStore interface {
	Publish(event event.Event) error
	Subscribe(h func(e event.Event)) (dispose func(), err error)
	Events() (events []event.Event, err error)
}

type EventBus interface {
	Publish(event event.Event) error
}

type Commands struct {
	sessions *session_repository.Repo
	projects *project_repository.Repo
	users    *user_repository.Repo
	store    EventStore
	bus      EventBus
	dispose  func()
}

func NewHandler(store EventStore, bus EventBus, sessions *session_repository.Repo,
	projects *project_repository.Repo,
	users *user_repository.Repo) *Commands {
	//sessionRepo, err := session_repository.New(store)
	//if err != nil {
	//	return nil, err
	//}
	//
	//userRepo, err := user_repository.New(store)
	//if err != nil {
	//	return nil, err
	//}
	//
	//projectRepo, err := project_repository.New(store)
	//if err != nil {
	//	return nil, err
	//}

	return &Commands{
		store:    store,
		bus:      bus,
		sessions: sessions,
		projects: projects,
		users:    users,
	}
}

func (s *Commands) Start() error {
	events, err := s.store.Events()
	if err != nil {
		return err
	}

	s.processEvents(events...)

	dispose, err := s.store.Subscribe(func(e event.Event) {
		s.processEvents(e)
	})
	if err != nil {
		return err
	}

	s.dispose = dispose

	return nil
}

func (s *Commands) processEvents(events ...event.Event) {
	for _, e := range events {
		switch e.AggregateType() {
		case event.User:
			s.users.ProcessEvents(e)
		case event.Project:
			s.projects.ProcessEvents(e)
		case event.Session:
			s.sessions.ProcessEvents(e)
		}
	}
}

func (s *Commands) publish(e event.Event) error {
	err := s.store.Publish(e)
	if err != nil {
		return err
	}

	return s.bus.Publish(e)
}
