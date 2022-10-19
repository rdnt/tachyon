package session_repository

import (
	"fmt"
	"sync"

	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/aggregate"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/event"
)

type EventStore interface {
	Events() ([]event.Event, error)
	Subscribe(h func(e event.Event)) (dispose func(), err error)
}

type Repo struct {
	store    EventStore
	mux      sync.Mutex
	sessions map[session.Id]*aggregate.Session
	dispose  func()
}

func (r *Repo) Session(id session.Id) (session.Session, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	s, ok := r.sessions[id]
	if !ok {
		return session.Session{}, command.ErrSessionNotFound
	}

	return s.Session, nil
}

func (r *Repo) ProjectSessionByName(pid project.Id, name string) (session.Session, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	for _, s := range r.sessions {
		if pid == s.ProjectId && s.Name == name {
			return s.Session, nil
		}
	}

	return session.Session{}, command.ErrSessionNotFound
}

func (r *Repo) ProjectSessions(pid project.Id) ([]session.Session, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	var sessions []session.Session
	for _, s := range r.sessions {
		if pid == s.ProjectId {
			sessions = append(sessions, s.Session)
		}
	}

	return sessions, nil
}

func (r *Repo) String() string {
	return fmt.Sprint(r.sessions)
}

func (r *Repo) processEvents(events ...event.Event) {
	r.mux.Lock()

	for _, e := range events {
		if e.AggregateType() != event.Session {
			continue
		}

		_, ok := r.sessions[session.Id(e.AggregateId())]
		if !ok {
			r.sessions[session.Id(e.AggregateId())] = &aggregate.Session{}
		}

		r.sessions[session.Id(e.AggregateId())].ProcessEvent(e)
	}

	r.mux.Unlock()
}

func New(store EventStore) (*Repo, error) {
	r := &Repo{
		store:    store,
		sessions: map[session.Id]*aggregate.Session{},
	}

	events, err := store.Events()
	if err != nil {
		return nil, err
	}

	r.processEvents(events...)

	dispose, err := store.Subscribe(func(e event.Event) {
		r.processEvents(e)
	})
	if err != nil {
		return nil, err
	}

	r.dispose = dispose

	return r, nil
}
