package event

import (
	"github.com/rdnt/tachyon/pkg/uuid"
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
