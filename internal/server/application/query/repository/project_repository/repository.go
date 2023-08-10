package project_repository

import (
	"fmt"
	"sync"

	"tachyon/internal/server/application/event"
	"tachyon/internal/server/application/query/aggregate"
	"tachyon/pkg/uuid"
)

type EventStore interface {
	Events() ([]event.Event, error)
	Subscribe(h func(e event.Event)) (dispose func(), err error)
}

type Repo struct {
	store    EventStore
	mux      sync.Mutex
	projects map[uuid.UUID]*aggregate.Project
	dispose  func()
}

func (r *Repo) String() string {
	return fmt.Sprint(r.projects)
}

func (r *Repo) processEvents(events ...event.Event) {
	r.mux.Lock()

	for _, e := range events {
		if e.AggregateType() != event.Project {
			continue
		}

		_, ok := r.projects[e.AggregateId()]
		if !ok {
			r.projects[e.AggregateId()] = aggregate.NewProject()
		}

		r.projects[e.AggregateId()].ProcessEvent(e)
	}

	r.mux.Unlock()
}

func New(store EventStore) (*Repo, error) {
	r := &Repo{
		store:    store,
		projects: map[uuid.UUID]*aggregate.Project{},
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
