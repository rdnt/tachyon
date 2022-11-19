package event

import (
	"tachyon/pkg/uuid"
)

const (
	ProjectCreated Type = "project_created"
)

type ProjectCreatedEvent struct {
	ProjectId uuid.UUID
	OwnerId   uuid.UUID
	Name      string
}

func (ProjectCreatedEvent) Type() Type {
	return ProjectCreated
}

func (ProjectCreatedEvent) AggregateType() AggregateType {
	return Project
}

func (e ProjectCreatedEvent) AggregateId() uuid.UUID {
	return e.ProjectId
}
