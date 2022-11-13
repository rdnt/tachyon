package event

import (
	"github.com/rdnt/tachyon/internal/client/application/domain/project"
	"github.com/rdnt/tachyon/pkg/uuid"
)

const DrawPixel Type = "draw_pixel"

type DrawPixelEvent struct {
	ProjectId uuid.UUID
	Color     project.Color
	Coords    project.Vector2
}

func (DrawPixelEvent) Type() Type {
	return DrawPixel
}
