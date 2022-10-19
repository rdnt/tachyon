package event

import (
	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/user"
)

const (
	PixelDrawn  Type = "pixel_drawn"
	PixelErased Type = "pixel_erased"
)

type PixelDrawnEvent struct {
	event

	UserId    user.Id
	ProjectId project.Id
	Color     project.Color
	Coords    project.Vector2
}

func NewPixelDrawnEvent(e PixelDrawnEvent) PixelDrawnEvent {
	e.typ = PathCreated
	e.aggregateType = Project
	e.aggregateId = uuid.UUID(e.ProjectId)

	return e
}

type PixelErasedEvent struct {
	event

	UserId    user.Id
	ProjectId project.Id
	Coords    project.Vector2
}

func NewPixelErasedEvent(e PixelErasedEvent) PixelErasedEvent {
	e.typ = PathCreated
	e.aggregateType = Project
	e.aggregateId = uuid.UUID(e.ProjectId)

	return e
}
