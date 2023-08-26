package project_repository

import (
	"errors"
	"sync"

	"tachyon/internal/server/application/command/aggregate"
	"tachyon/internal/server/application/domain/project"
	"tachyon/internal/server/application/event"
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
}

func New(store EventStore) (*Repo, error) {
	r := &Repo{
		store:    store,
		projects: map[uuid.UUID]*aggregate.Project{},
	}

	return r, nil
}

func (r *Repo) ProcessEvents(events ...event.Event) {
	r.mux.Lock()

	for _, e := range events {
		_, ok := r.projects[e.AggregateId()]
		if !ok {
			r.projects[e.AggregateId()] = &aggregate.Project{}
		}

		r.projects[e.AggregateId()].ProcessEvent(e)
	}

	r.mux.Unlock()
}

var ErrProjectNotFound = errors.New("project not found")

func (r *Repo) Project(id uuid.UUID) (project.Project, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	p, ok := r.projects[id]
	if !ok {
		return project.Project{}, ErrProjectNotFound
	}

	return p.Project, nil
}

func (r *Repo) UserProjectByName(uid uuid.UUID, name string) (project.Project, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	for _, p := range r.projects {
		if uid == p.OwnerId && p.Name == name {
			return p.Project, nil
		}
	}

	return project.Project{}, ErrProjectNotFound
}
