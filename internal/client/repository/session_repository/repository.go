package session_repository

import (
	"fmt"
	"sync"

	"github.com/samber/lo"

	"tachyon/internal/client/application"
	"tachyon/internal/client/application/aggregate"
	"tachyon/internal/client/application/domain/session"
	"tachyon/internal/client/remote"
	"tachyon/internal/pkg/event"
	"tachyon/pkg/uuid"
)

type Repo struct {
	mux       sync.Mutex
	sessions  map[uuid.UUID]*aggregate.Session
	sessionId uuid.UUID
}

func (r *Repo) Session() (session.Session, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	if r.sessionId == uuid.Nil {
		return session.Session{}, application.ErrSessionNotFound
	}

	s, ok := r.sessions[r.sessionId]
	if !ok {
		return session.Session{}, application.ErrSessionNotFound
	}

	return s.Session, nil
}

func (r *Repo) Sessions() ([]session.Session, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	return lo.Map(lo.Values(r.sessions), func(item *aggregate.Session, index int) session.Session {
		return item.Session
	}), nil
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

		if e.Type() == event.SessionCreated {
			r.sessionId = uuid.MustParse(e.AggregateId())
		}
	}

	r.mux.Unlock()
}

func New() (*Repo, error) {
	r := &Repo{
		sessions: map[uuid.UUID]*aggregate.Session{},
	}

	return r, nil
}
