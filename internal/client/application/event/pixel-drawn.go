package event

import (
	"github.com/rdnt/tachyon/internal/client/application/domain/project"
	"github.com/rdnt/tachyon/pkg/uuid"
)

const PixelDrawn Type = "pixel_drawn"

type PixelDrawnEvent struct {
	UserId    uuid.UUID
	ProjectId uuid.UUID
	Color     project.Color
	Coords    project.Vector2
}

func (PixelDrawnEvent) Type() Type {
	return PixelDrawn
}
