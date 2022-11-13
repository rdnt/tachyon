package event

import (
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/pkg/uuid"
)

const ErasePixel Type = "erase_pixel"

type ErasePixelEvent struct {
	ProjectId uuid.UUID
	Coords    project.Vector2
}

func (ErasePixelEvent) Type() Type {
	return ErasePixel
}
