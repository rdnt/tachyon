package redisevent

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type PixelDrawnEvent struct {
	UserId    string     `json:"userId"`
	ProjectId string     `json:"projectId"`
	Color     string     `json:"color"`
	Coords    IntVector2 `json:"coords"`
}

func PixelDrawnEventToJSON(e event.PixelDrawnEvent) ([]byte, error) {
	evt := PixelDrawnEvent{
		UserId:    uuid.UUID(e.UserId).String(),
		ProjectId: uuid.UUID(e.ProjectId).String(),
		Color:     e.Color.String(),
		Coords: IntVector2{
			X: e.Coords.X,
			Y: e.Coords.Y,
		},
	}

	return json.Marshal(evt)
}

func PixelDrawnEventFromJSON(b []byte) (event.PixelDrawnEvent, error) {
	var evt PixelDrawnEvent
	err := json.Unmarshal(b, &evt)
	if err != nil {
		return event.PixelDrawnEvent{}, err
	}

	uid, err := uuid.Parse(evt.UserId)
	if err != nil {
		return event.PixelDrawnEvent{}, err
	}

	pid, err := uuid.Parse(evt.ProjectId)
	if err != nil {
		return event.PixelDrawnEvent{}, err
	}

	clr, err := project.ColorFromString(evt.Color)
	if err != nil {
		return event.PixelDrawnEvent{}, err
	}

	return event.PixelDrawnEvent{
		UserId:    uuid.UUID(uid),
		ProjectId: uuid.UUID(pid),
		Color:     clr,
		Coords: project.Vector2{
			X: evt.Coords.X,
			Y: evt.Coords.Y,
		},
	}, nil
}
