package event

import (
	"tachyon/internal/server/application/domain/project"
	"tachyon/pkg/uuid"
)

const (
	PixelDrawn  Type = "pixel_drawn"
	PixelErased Type = "pixel_erased"
)

type PixelDrawnEvent struct {
	UserId    uuid.UUID
	ProjectId uuid.UUID
	Color     project.Color
	Coords    project.Vector2
}

func (PixelDrawnEvent) Type() Type {
	return PixelDrawn
}

func (PixelDrawnEvent) AggregateType() AggregateType {
	return Project
}

func (e PixelDrawnEvent) AggregateId() uuid.UUID {
	return e.ProjectId
}

type PixelErasedEvent struct {
	UserId    uuid.UUID
	ProjectId uuid.UUID
	Coords    project.Vector2
}

func (PixelErasedEvent) Type() Type {
	return PixelErased
}

func (PixelErasedEvent) AggregateType() AggregateType {
	return Project
}

func (e PixelErasedEvent) AggregateId() uuid.UUID {
	return e.ProjectId
}
