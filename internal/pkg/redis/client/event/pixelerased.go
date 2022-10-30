package event

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type PixelErasedEvent struct {
	UserId    string     `json:"userId"`
	ProjectId string     `json:"projectId"`
	Coords    IntVector2 `json:"coords"`
}

func PixelErasedEventToJSON(e event.PixelErasedEvent) ([]byte, error) {
	evt := PixelErasedEvent{
		UserId:    e.UserId.String(),
		ProjectId: e.ProjectId.String(),
		Coords: IntVector2{
			X: e.Coords.X,
			Y: e.Coords.Y,
		},
	}

	return json.Marshal(evt)
}

func PixelErasedEventFromJSON(b []byte) (event.PixelErasedEvent, error) {
	var evt PixelErasedEvent
	err := json.Unmarshal(b, &evt)
	if err != nil {
		return event.PixelErasedEvent{}, err
	}

	uid, err := uuid.Parse(evt.UserId)
	if err != nil {
		return event.PixelErasedEvent{}, err
	}

	pid, err := uuid.Parse(evt.ProjectId)
	if err != nil {
		return event.PixelErasedEvent{}, err
	}

	return event.PixelErasedEvent{
		UserId:    uid,
		ProjectId: pid,
		Coords: project.Vector2{
			X: evt.Coords.X,
			Y: evt.Coords.Y,
		},
	}, nil
}
