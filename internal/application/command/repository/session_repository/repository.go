package session_repository

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/event"
)

type EventStore interface {
	Events() ([]event.Event, error)
}

type Repo struct {
	events   EventStore
	sessions []session.Session
}

func (r *Repo) Session(id session.Id) (session.Session, error) {
	panic("not implemented")
}

func (r *Repo) String() string {
	b, err := json.Marshal(r.sessions)
	if err != nil {
		return "error"
	}

	return string(b)
}

func New(events EventStore) *Repo {
	return &Repo{
		events:   events,
		sessions: []session.Session{},
	}
}
