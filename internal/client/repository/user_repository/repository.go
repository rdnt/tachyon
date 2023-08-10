package user_repository

import (
	"fmt"
	"sync"

	"github.com/samber/lo"

	"tachyon/internal/client/application"
	"tachyon/internal/client/application/aggregate"
	"tachyon/internal/client/application/domain/user"
	"tachyon/internal/client/remote"
	"tachyon/internal/pkg/event"
	"tachyon/pkg/uuid"
)

type Repo struct {
	mux    sync.Mutex
	users  map[uuid.UUID]*aggregate.User
	userId uuid.UUID
}

func (r *Repo) User() (user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	if r.userId == uuid.Nil {
		return user.User{}, application.ErrUserNotFound
	}

	u, ok := r.users[r.userId]
	if !ok {
		return user.User{}, application.ErrUserNotFound
	}

	return u.User, nil
}

func (r *Repo) Users() ([]user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	return lo.Map(lo.Values(r.users), func(item *aggregate.User, index int) user.User {
		return item.User
	}), nil
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

		if e.Type() == event.Connected {
			r.userId = uuid.MustParse(e.AggregateId())
		}
	}

	r.mux.Unlock()
}

func New() (*Repo, error) {
	r := &Repo{
		users: map[uuid.UUID]*aggregate.User{},
	}

	return r, nil
}
