package sessionrepo

import (
	"errors"
	"fmt"
	"sync"

	"tachyon2/pkg/logger"
	"tachyon2/src/application/domain/session"
)

var ErrNotFound = errors.New("session not found")
var ErrExists = errors.New("session already exists")

type Repository struct {
	log         *logger.Logger
	sessionsMux sync.Mutex
	sessions    map[session.Id]session.Session
}

func (r *Repository) CreateSession(s session.Session) (session.Session, error) {
	r.sessionsMux.Lock()
	defer r.sessionsMux.Unlock()

	_, ok := r.sessions[s.Id]
	if ok {
		return session.Session{}, ErrExists
	}

	r.sessions[s.Id] = s

	r.log.Println("session created:", s)
	r.log.Println(r)

	return s, nil
}

func (r *Repository) Session(id session.Id) (session.Session, error) {
	r.sessionsMux.Lock()
	defer r.sessionsMux.Unlock()

	s, ok := r.sessions[id]
	if !ok {
		return session.Session{}, ErrNotFound
	}

	return s, nil
}

func (r *Repository) Sessions() (map[session.Id]session.Session, error) {
	r.sessionsMux.Lock()
	defer r.sessionsMux.Unlock()

	return r.sessions, nil
}

func (r *Repository) UpdateSession(s session.Session) (session.Session, error) {
	r.sessionsMux.Lock()
	defer r.sessionsMux.Unlock()

	_, ok := r.sessions[s.Id]
	if !ok {
		return session.Session{}, ErrNotFound
	}

	r.sessions[s.Id] = s
	return s, nil
}

func (r *Repository) DeleteSession(id session.Id) error {
	r.sessionsMux.Lock()
	defer r.sessionsMux.Unlock()

	delete(r.sessions, id)

	r.log.Println("session deleted:", id)
	r.log.Println(r)

	return nil
}

func (r *Repository) String() string {
	return fmt.Sprintf("=== %v", r.sessions)
}

func New() *Repository {
	return &Repository{
		sessions: map[session.Id]session.Session{},
		log:      logger.New("sessions", logger.RedFg),
	}
}
