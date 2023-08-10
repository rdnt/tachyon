package project_repository

import (
	"fmt"
	"sync"

	"github.com/samber/lo"

	"tachyon/internal/client/application"
	"tachyon/internal/client/application/aggregate"
	"tachyon/internal/client/application/domain/project"
	"tachyon/internal/client/remote"
	"tachyon/internal/pkg/event"
	"tachyon/pkg/uuid"
)

type Repo struct {
	mux       sync.Mutex
	projects  map[uuid.UUID]*aggregate.Project
	projectId uuid.UUID
}

func (r *Repo) Project() (project.Project, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	if r.projectId == uuid.Nil {
		return project.Project{}, application.ErrProjectNotFound
	}

	p, ok := r.projects[r.projectId]
	if !ok {
		return project.Project{}, application.ErrProjectNotFound
	}

	return p.Project, nil
}

func (r *Repo) Projects() ([]project.Project, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	return lo.Map(lo.Values(r.projects), func(item *aggregate.Project, index int) project.Project {
		return item.Project
	}), nil
}

func (r *Repo) String() string {
	return fmt.Sprint(r.projects)
}

func (r *Repo) ProcessEvents(events ...remote.Event) {
	r.mux.Lock()

	for _, e := range events {
		if e.AggregateType() != event.Project {
			continue
		}

		_, ok := r.projects[uuid.MustParse(e.AggregateId())]
		if !ok {
			r.projects[uuid.MustParse(e.AggregateId())] = &aggregate.Project{}
		}

		r.projects[uuid.MustParse(e.AggregateId())].ProcessEvent(e)

		if e.Type() == event.ProjectCreated {
			r.projectId = uuid.MustParse(e.AggregateId())
		}
	}

	r.mux.Unlock()
}

func New() (*Repo, error) {
	r := &Repo{
		projects: map[uuid.UUID]*aggregate.Project{},
	}

	return r, nil
}
