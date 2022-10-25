package router

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/application/event"
)

type Event struct {
	Type          string `json:"type"`
	AggregateType string `json:"aggregateType"`
	AggregateId   string `json:"aggregateId"`
}

//const (
//	PixelDrawn  event.Type = "pixel_drawn"
//	PixelErased event.Type = "pixel_erased"
//)

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
		UserId:    e.UserId.String(),
		ProjectId: e.ProjectId.String(),
		Color:     e.Color.String(),
		X:         e.Coords.X,
		Y:         e.Coords.Y,
	}

	return json.Marshal(evt)
}

type PixelErasedEvent struct {
	Event

	UserId    string `json:"userId"`
	ProjectId string `json:"projectId"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
}
