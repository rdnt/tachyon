package user_repository

import (
	"fmt"
	"sync"

	"tachyon/internal/client/application"
	"tachyon/internal/client/application/aggregate"
	"tachyon/internal/client/application/domain/user"
	"tachyon/internal/client/remote"
	"tachyon/internal/pkg/event"
	"tachyon/pkg/uuid"
)

type Repo struct {
	mux   sync.Mutex
	users map[uuid.UUID]*aggregate.User
}

func (r *Repo) User(id uuid.UUID) (user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	u, ok := r.users[id]
	if !ok {
		return user.User{}, application.ErrUserNotFound
	}

	return u.User, nil
}

func (r *Repo) String() string {
	return fmt.Sprint(r.users)
}

func (r *Repo) ProcessEvents(events ...remote.Event) {
	r.mux.Lock()

	for _, e := range events {
		if e.AggregateType() != event.User {
			continue
		}

		_, ok := r.users[uuid.MustParse(e.AggregateId())]
		if !ok {
			r.users[uuid.MustParse(e.AggregateId())] = &aggregate.User{}
		}

		r.users[uuid.MustParse(e.AggregateId())].ProcessEvent(e)
	}

	r.mux.Unlock()
}

func New() (*Repo, error) {
	r := &Repo{
		users: map[uuid.UUID]*aggregate.User{},
	}

	return r, nil
}
