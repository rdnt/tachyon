package event

import (
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/project/path"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/pkg/uuid"
)

const (
	PathCreated Type = "path_created"
)

type PathCreatedEvent struct {
	event

	PathId    path.Id
	UserId    user.Id
	ProjectId project.Id
	Tool      path.Tool
	Color     path.Color
	Point     path.Vector2
}

func NewPathCreatedEvent(e PathCreatedEvent) PathCreatedEvent {
	e.typ = PathCreated
	e.aggregateType = Project
	e.aggregateId = uuid.UUID(e.ProjectId)

	return e
}
