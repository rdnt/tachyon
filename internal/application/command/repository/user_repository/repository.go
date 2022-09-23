package user_repository

import (
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
)

type EventStore interface {
	Events() ([]event.Event, error)
}

type Repo struct {
	events EventStore
	//users  []user.User
}

func (r *Repo) User(id user.Id) (user.User, error) {
	events, err := r.events.Events()
	if err != nil {
		return user.User{}, err
	}

	var evts []event.Event
	for _, e := range events {
		if e.AggregateType() != event.User || e.AggregateId() != string(id) {
			continue
		}

		evts = append(evts, e)
	}

	return newUserFromEvents(evts), nil
}

func newUserFromEvents(events []event.Event) user.User {
	var u user.User

	for _, e := range events {
		switch e := e.(type) {
		case event.UserCreatedEvent:
			u.Id = e.Id
			u.Name = e.Name
		default:
			continue
		}
	}

	return u
}

//func (r *Repo) String() string {
//	b, err := json.Marshal(r.users)
//	if err != nil {
//		return "error"
//	}
//
//	return string(b)
//}

func New(events EventStore) *Repo {
	return &Repo{
		events: events,
		//users:  []user.User{},
	}
}
