package session_repository

import (
	"errors"
	"sync"

	"tachyon/internal/server/application/command/aggregate"
	"tachyon/internal/server/application/domain/session"
	"tachyon/internal/server/application/event"
	"tachyon/pkg/uuid"
)

type EventStore interface {
	Events() ([]event.Event, error)
	Subscribe(handler func(e event.Event)) (dispose func(), err error)
}

type Repo struct {
	store    EventStore
	mux      sync.Mutex
	sessions map[uuid.UUID]*aggregate.Session
}

func New(store EventStore) (*Repo, error) {
	r := &Repo{
		store:    store,
		sessions: map[uuid.UUID]*aggregate.Session{},
	}

	return r, nil
}

func (r *Repo) ProcessEvents(events ...event.Event) {
	r.mux.Lock()

	for _, e := range events {
		_, ok := r.sessions[e.AggregateId()]
		if !ok {
			r.sessions[e.AggregateId()] = &aggregate.Session{}
		}

		r.sessions[e.AggregateId()].ProcessEvent(e)
	}

	r.mux.Unlock()
}

var ErrSessionNotFound = errors.New("session not found")

func (r *Repo) Session(id uuid.UUID) (session.Session, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	s, ok := r.sessions[id]
	if !ok {
		return session.Session{}, ErrSessionNotFound
	}

	return s.Session, nil
}

func (r *Repo) ProjectSessionByName(pid uuid.UUID, name string) (session.Session, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	for _, s := range r.sessions {
		if pid == s.ProjectId && s.Name == name {
			return s.Session, nil
		}
	}

	return session.Session{}, ErrSessionNotFound
}

func (r *Repo) ProjectSessions(pid uuid.UUID) ([]session.Session, error) {
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
