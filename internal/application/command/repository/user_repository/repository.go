package user_repository

import (
	"fmt"
	"sync"

	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/aggregate"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
)

type EventStore interface {
	Events() ([]event.EventIface, error)
	Subscribe(h func(e event.EventIface)) (dispose func(), err error)
}

type Repo struct {
	store   EventStore
	mux     sync.Mutex
	users   map[user.Id]*aggregate.User
	dispose func()
}

func (r *Repo) User(id user.Id) (user.User, error) {
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

func (r *Repo) processEvents(events ...event.EventIface) {
	r.mux.Lock()

	for _, e := range events {
		if e.AggregateType() != event.User {
			continue
		}

		_, ok := r.users[user.Id(e.AggregateId())]
		if !ok {
			r.users[user.Id(e.AggregateId())] = &aggregate.User{}
		}

		r.users[user.Id(e.AggregateId())].ProcessEvent(e)
	}

	r.mux.Unlock()
}

func New(store EventStore) (*Repo, error) {
	r := &Repo{
		store: store,
		users: map[user.Id]*aggregate.User{},
	}

	events, err := store.Events()
	if err != nil {
		return nil, err
	}

	r.processEvents(events...)

	dispose, err := store.Subscribe(func(e event.EventIface) {
		r.processEvents(e)
	})
	if err != nil {
		return nil, err
	}

	r.dispose = dispose

	return r, nil
}
