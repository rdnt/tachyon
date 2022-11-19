package event

import "github.com/rdnt/tachyon/pkg/uuid"

const (
	CreateProject Type = "create_project"
)

type CreateProjectEvent struct {
	ProjectId uuid.UUID
	Name      string
}

func (CreateProjectEvent) Type() Type {
	return CreateProject
}

func (CreateProjectEvent) AggregateType() Type {
	return Project
}

func (e CreateProjectEvent) AggregateId() uuid.UUID {
	return e.ProjectId
}
