package project_repository

import (
	"fmt"
	"sync"

	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/aggregate"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
)

type EventStore interface {
	Events() ([]event.Event, error)
	Subscribe(h func(e event.Event)) (dispose func(), err error)
}

type Repo struct {
	store    EventStore
	mux      sync.Mutex
	projects map[project.Id]*aggregate.Project
	dispose  func()
}

func (r *Repo) Project(id project.Id) (project.Project, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	p, ok := r.projects[id]
	if !ok {
		return project.Project{}, command.ErrProjectNotFound
	}

	return p.Project, nil
}

func (r *Repo) UserProjectByName(uid user.Id, name string) (project.Project, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	for _, p := range r.projects {
		if uid == p.OwnerId && p.Name == name {
			return p.Project, nil
		}
	}

	return project.Project{}, command.ErrProjectNotFound
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

		_, ok := r.projects[project.Id(e.AggregateId())]
		if !ok {
			r.projects[project.Id(e.AggregateId())] = &aggregate.Project{}
		}

		r.projects[project.Id(e.AggregateId())].ProcessEvent(e)
	}

	r.mux.Unlock()
}

func New(store EventStore) (*Repo, error) {
	r := &Repo{
		store:    store,
		projects: map[project.Id]*aggregate.Project{},
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
