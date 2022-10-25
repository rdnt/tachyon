package router

import (
	"encoding/json"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/pkg/uuid"

	"github.com/rdnt/tachyon/internal/application/event"
)

type Event struct {
	Type          string `json:"type"`
	AggregateType string `json:"aggregateType"`
	AggregateId   string `json:"aggregateId"`
}

// const (
//	PixelDrawn  event.Type = "pixel_drawn"
//	PixelErased event.Type = "pixel_erased"
// )

type PixelDrawnEvent struct {
	Event

	UserId    string `json:"userId"`
	ProjectId string `json:"projectId"`
	Color     string `json:"color"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
}

func jsonEvent(e event.Event) Event {
	return Event{
		Type:          e.Type(),
		AggregateType: string(e.AggregateType()),
		AggregateId:   e.AggregateId().String(),
	}
}

func PixelDrawnEventJSON(e event.PixelDrawnEvent) ([]byte, error) {
	evt := PixelDrawnEvent{
		Event:     jsonEvent(e),
		UserId:    uuid.UUID(e.UserId).String(),
		ProjectId: uuid.UUID(e.ProjectId).String(),
		Color:     e.Color.String(),
		X:         e.Coords.X,
		Y:         e.Coords.Y,
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

	return event.NewPixelDrawnEvent(event.PixelDrawnEvent{
		UserId:    user.Id(uuid.Parse(evt.UserId)),
		ProjectId: project.Id{},
		Color:     project.Color{},
		Coords:    project.Vector2{},
	})
}

type PixelErasedEvent struct {
	Event

	UserId    string `json:"userId"`
	ProjectId string `json:"projectId"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
}
