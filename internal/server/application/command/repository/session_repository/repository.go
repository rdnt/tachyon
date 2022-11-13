package session_repository

import (
	"fmt"
	"sync"

	"github.com/rdnt/tachyon/internal/server/application/command"
	"github.com/rdnt/tachyon/internal/server/application/command/aggregate"
	"github.com/rdnt/tachyon/internal/server/application/domain/session"
	"github.com/rdnt/tachyon/internal/server/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type EventStore interface {
	Events() ([]event.Event, error)
	Subscribe(handler func(e event.Event)) (dispose func(), err error)
}

type Repo struct {
	store    EventStore
	mux      sync.Mutex
	sessions map[uuid.UUID]*aggregate.Session
	dispose  func()
}

func (r *Repo) Session(id uuid.UUID) (session.Session, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	s, ok := r.sessions[id]
	if !ok {
		return session.Session{}, command.ErrSessionNotFound
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

	return session.Session{}, command.ErrSessionNotFound
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

func (r *Repo) String() string {
	return fmt.Sprint(r.sessions)
}

func (r *Repo) processEvents(events ...event.Event) {
	r.mux.Lock()

	for _, e := range events {
		if e.AggregateType() != event.Session {
			continue
		}

		_, ok := r.sessions[uuid.UUID(e.AggregateId())]
		if !ok {
			r.sessions[uuid.UUID(e.AggregateId())] = &aggregate.Session{}
		}

		r.sessions[uuid.UUID(e.AggregateId())].ProcessEvent(e)
	}

	r.mux.Unlock()
}

func New(store EventStore) (*Repo, error) {
	r := &Repo{
		store:    store,
		sessions: map[uuid.UUID]*aggregate.Session{},
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
