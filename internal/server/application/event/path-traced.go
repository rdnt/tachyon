package event

import (
	"tachyon/internal/server/application/domain/project/path"
	"tachyon/pkg/uuid"
)

const (
	PathTraced Type = "path-traced"
)

type PathTracedEvent struct {
	PathId    uuid.UUID
	UserId    uuid.UUID
	ProjectId uuid.UUID
	Point     path.Vector2
}

func (PathTracedEvent) Type() Type {
	return PathTraced
}

func (PathTracedEvent) AggregateType() AggregateType {
	return Project
}

func (e PathTracedEvent) AggregateId() uuid.UUID {
	return uuid.UUID(e.ProjectId)
}
