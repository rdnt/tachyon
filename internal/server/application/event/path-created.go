package event

import (
	"tachyon/internal/server/application/domain/project/path"
	"tachyon/pkg/uuid"
)

const (
	PathCreated Type = "path-created"
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
