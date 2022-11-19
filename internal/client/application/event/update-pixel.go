package event

import (
	"github.com/rdnt/tachyon/internal/client/application/domain/project"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type UpdatePixelEvent struct {
	ProjectId uuid.UUID
	Color     project.Color
	Coords    project.Vector2
}

func (UpdatePixelEvent) Type() Type {
	return UpdatePixel
}
