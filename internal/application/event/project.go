package event

import (
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/pkg/uuid"
)

const (
	ProjectCreated Type = "project_created"
)

type ProjectCreatedEvent struct {
	event

	Id      project.Id
	OwnerId user.Id
	Name    string
}

func NewProjectCreatedEvent(e ProjectCreatedEvent) ProjectCreatedEvent {
	e.typ = ProjectCreated
	e.aggregateType = Project
	e.aggregateId = uuid.UUID(e.Id)

	return e
}
