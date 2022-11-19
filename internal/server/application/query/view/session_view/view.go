package session_view

import (
	"encoding/json"
	"errors"

	"tachyon/internal/server/application/domain/session"
	"tachyon/pkg/uuid"
)

var ErrSessionNotFound = errors.New("session not found")

type View struct {
	sessions map[uuid.UUID]session.Session
}

func (v *View) Session(id uuid.UUID) (session.Session, error) {
	s, ok := v.sessions[id]
	if !ok {
		return session.Session{}, ErrSessionNotFound
	}

	return s, nil
}

func (v *View) CreateSession(s session.Session) error {
	v.sessions[s.Id] = s
	return nil
}

func (v *View) String() string {
	b, err := json.Marshal(v.sessions)
	if err != nil {
		return "error"
	}

	return string(b)
}

func New() *View {
	r := &View{
		sessions: map[uuid.UUID]session.Session{},
	}

	return r
}
