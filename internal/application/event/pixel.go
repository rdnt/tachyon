package event

import (
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/pkg/uuid"
)

const (
	PixelDrawn  Type = "pixel_drawn"
	PixelErased Type = "pixel_erased"
)

type PixelDrawnEvent struct {
	event

	UserId    string
	ProjectId string
	Color     string
	X         string `json:"x"`
}

func NewPixelDrawnEvent(e PixelDrawnEvent) PixelDrawnEvent {
	e.typ = PixelDrawn
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
	e.typ = PixelErased
	e.aggregateType = Project
	e.aggregateId = uuid.UUID(e.ProjectId)

	return e
}
