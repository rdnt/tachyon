package event

import (
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/pkg/uuid"
)

const PixelErased Type = "pixel_erased"

type PixelErasedEvent struct {
	UserId    uuid.UUID
	ProjectId uuid.UUID
	Coords    project.Vector2
}

func (PixelErasedEvent) Type() Type {
	return PixelErased
}
