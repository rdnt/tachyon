package user_repository

import (
	"errors"
	"sync"

	"tachyon/internal/server/application/command/aggregate"
	"tachyon/internal/server/application/domain/user"
	"tachyon/internal/server/application/event"
	"tachyon/pkg/uuid"
)

type EventStore interface {
	//Events() ([]event.Event, error)
	Subscribe(h func(e event.Event)) (dispose func(), err error)
}

type Repo struct {
	store EventStore
	mux   sync.Mutex
	users map[uuid.UUID]*aggregate.User
}

func New(store EventStore) (*Repo, error) {
	r := &Repo{
		store: store,
		users: map[uuid.UUID]*aggregate.User{},
	}

	return r, nil
}

func (r *Repo) ProcessEvents(events ...event.Event) {
	r.mux.Lock()

	for _, e := range events {
		_, ok := r.users[e.AggregateId()]
		if !ok {
			r.users[e.AggregateId()] = &aggregate.User{}
		}

		r.users[e.AggregateId()].ProcessEvent(e)
	}

	r.mux.Unlock()
}

var ErrUserNotFound = errors.New("user not found")

func (r *Repo) User(id uuid.UUID) (user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	u, ok := r.users[id]
	if !ok {
		return user.User{}, ErrUserNotFound
	}

	return u.User, nil
}

func (r *Repo) UserByName(name string) (user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	for _, u := range r.users {
		if u.Name == name {
			return u.User, nil
		}
	}

	return user.User{}, ErrUserNotFound
}
