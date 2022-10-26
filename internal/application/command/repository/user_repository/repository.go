package user_repository

import (
	"fmt"
	"sync"

	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/aggregate"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type EventStore interface {
	Events() ([]event.Event, error)
	Subscribe(h func(e event.Event)) (dispose func(), err error)
}

type Repo struct {
	store   EventStore
	mux     sync.Mutex
	users   map[uuid.UUID]*aggregate.User
	dispose func()
}

func (r *Repo) User(id uuid.UUID) (user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	u, ok := r.users[id]
	if !ok {
		return user.User{}, command.ErrUserNotFound
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

	return user.User{}, command.ErrUserNotFound
}

func (r *Repo) String() string {
	return fmt.Sprint(r.users)
}

func (r *Repo) processEvents(events ...event.Event) {
	r.mux.Lock()

	for _, e := range events {
		if e.AggregateType() != event.User {
			continue
		}

		_, ok := r.users[uuid.UUID(e.AggregateId())]
		if !ok {
			r.users[uuid.UUID(e.AggregateId())] = &aggregate.User{}
		}

		r.users[uuid.UUID(e.AggregateId())].ProcessEvent(e)
	}

	r.mux.Unlock()
}

func New(store EventStore) (*Repo, error) {
	r := &Repo{
		store: store,
		users: map[uuid.UUID]*aggregate.User{},
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
