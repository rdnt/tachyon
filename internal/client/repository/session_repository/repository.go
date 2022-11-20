package session_repository

import (
	"fmt"
	"sync"

	"tachyon/internal/client/application"
	"tachyon/internal/client/application/aggregate"
	"tachyon/internal/client/application/domain/session"
	"tachyon/internal/client/remote"
	"tachyon/internal/pkg/event"
	"tachyon/pkg/uuid"
)

type Repo struct {
	mux      sync.Mutex
	sessions map[uuid.UUID]*aggregate.Session
}

func (r *Repo) Session(id uuid.UUID) (session.Session, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	s, ok := r.sessions[id]
	if !ok {
		return session.Session{}, application.ErrSessionNotFound
	}

	return s.Session, nil
}

func (r *Repo) String() string {
	return fmt.Sprint(r.sessions)
}

func (r *Repo) ProcessEvents(events ...remote.Event) {
	r.mux.Lock()

	for _, e := range events {
		if e.AggregateType() != event.Session {
			continue
		}

		_, ok := r.sessions[uuid.MustParse(e.AggregateId())]
		if !ok {
			r.sessions[uuid.MustParse(e.AggregateId())] = &aggregate.Session{}
		}

		r.sessions[uuid.MustParse(e.AggregateId())].ProcessEvent(e)
	}

	r.mux.Unlock()
}

func New() (*Repo, error) {
	r := &Repo{
		sessions: map[uuid.UUID]*aggregate.Session{},
	}

	return r, nil
}
