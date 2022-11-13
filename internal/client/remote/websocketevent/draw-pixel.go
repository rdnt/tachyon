package websocketevent

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/client/application/domain/project"
	"github.com/rdnt/tachyon/internal/client/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type DrawPixelEvent struct {
	Event     string     `json:"event"`
	ProjectId string     `json:"projectId"`
	Color     string     `json:"color"`
	Coords    IntVector2 `json:"coords"`
}

func DrawPixelEventFromJSON(b []byte) (event.DrawPixelEvent, error) {
	var e DrawPixelEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return event.DrawPixelEvent{}, err
	}

	pid, err := uuid.Parse(e.ProjectId)
	if err != nil {
		return event.DrawPixelEvent{}, err
	}

	color, err := project.ColorFromString(e.Color)
	if err != nil {
		return event.DrawPixelEvent{}, err
	}

	return event.DrawPixelEvent{
		ProjectId: pid,
		Color:     color,
		Coords: project.Vector2{
			X: e.Coords.X,
			Y: e.Coords.Y,
		},
	}, nil
}

func DrawPixelEventToJSON(e event.DrawPixelEvent) ([]byte, error) {
	e2 := DrawPixelEvent{
		Event:     e.Type().String(),
		ProjectId: e.ProjectId.String(),
		Color:     e.Color.String(),
		Coords: IntVector2{
			X: e.Coords.X,
			Y: e.Coords.Y,
		},
	}

	return json.Marshal(e2)
}