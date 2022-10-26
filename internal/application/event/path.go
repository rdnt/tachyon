package event

import (
	"github.com/rdnt/tachyon/internal/application/domain/project/path"
	"github.com/rdnt/tachyon/pkg/uuid"
)

const (
	PathCreated Type = "path_created"
)

type PathCreatedEvent struct {
	PathId    uuid.UUID
	UserId    uuid.UUID
	ProjectId uuid.UUID
	Tool      path.Tool
	Color     path.Color
	Point     path.Vector2
}

func (PathCreatedEvent) Type() Type {
	return PathCreated
}

func (PathCreatedEvent) AggregateType() AggregateType {
	return Project
}

func (e PathCreatedEvent) AggregateId() uuid.UUID {
	return uuid.UUID(e.ProjectId)
}
