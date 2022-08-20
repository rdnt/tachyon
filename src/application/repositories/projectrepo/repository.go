package projectrepo

import (
	"errors"
	"fmt"
	"sync"

	"tachyon2/pkg/logger"
	"tachyon2/src/application/domain/project"
)

var ErrNotFound = errors.New("project not found")
var ErrExists = errors.New("project already exists")

type Repository struct {
	log      *logger.Logger
	mux      sync.Mutex
	projects map[project.Id]project.Project
}

func (r *Repository) CreateProject(p project.Project) (project.Project, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	_, ok := r.projects[p.Id]
	if ok {
		return project.Project{}, ErrExists
	}

	r.projects[p.Id] = p

	r.log.Println("project created:", p)
	r.log.Println(r)

	return p, nil
}

func (r *Repository) Project(id project.Id) (project.Project, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	s, ok := r.projects[id]
	if !ok {
		return project.Project{}, ErrNotFound
	}

	return s, nil
}

func (r *Repository) Projects() (map[project.Id]project.Project, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.projects, nil
}

func (r *Repository) UpdateProject(p project.Project) (project.Project, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	_, ok := r.projects[p.Id]
	if !ok {
		return project.Project{}, ErrNotFound
	}

	r.projects[p.Id] = p
	return p, nil
}

func (r *Repository) DeleteProject(id project.Id) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	delete(r.projects, id)

	r.log.Println("project deleted:", id)
	r.log.Println(r)

	return nil
}

func (r *Repository) String() string {
	return fmt.Sprintf("=== %v", r.projects)
}

func New() *Repository {
	return &Repository{
		projects: map[project.Id]project.Project{},
		log:      logger.New("projects", logger.RedFg),
	}
}
