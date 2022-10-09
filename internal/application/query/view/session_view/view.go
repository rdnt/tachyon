package session_view

import (
	"encoding/json"
	"errors"

	"github.com/rdnt/tachyon/internal/application/domain/session"
)

var ErrSessionNotFound = errors.New("session not found")

type View struct {
	sessions map[session.Id]session.Session
}

func (v *View) Session(id session.Id) (session.Session, error) {
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
		sessions: map[session.Id]session.Session{},
	}

	return r
}
