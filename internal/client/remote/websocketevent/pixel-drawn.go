package websocketevent

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/client/application/event"
)

type PixelDrawnEvent struct {
	Event     string     `json:"event"`
	UserId    string     `json:"userId"`
	ProjectId string     `json:"projectId"`
	Color     string     `json:"color"`
	Coords    IntVector2 `json:"coords"`
}

func PixelDrawnEventToJSON(e event.PixelDrawnEvent) ([]byte, error) {
	e2 := PixelDrawnEvent{
		Event:     event.PixelUpdated.String(),
		UserId:    e.UserId.String(),
		ProjectId: e.ProjectId.String(),
		Color:     e.Color.String(),
		Coords: IntVector2{
			X: e.Coords.X,
			Y: e.Coords.Y,
		},
	}

	return json.Marshal(e2)
}
