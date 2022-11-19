package websocketevent

import (
	"encoding/json"
	"errors"
	"fmt"

	"tachyon/internal/client/application/event"
)

type Event struct {
	Event string `json:"event"`
}

func FromJSON(b []byte) (event.Event, error) {
	var evt Event
	err := json.Unmarshal(b, &evt)
	if err != nil {
		return nil, err
	}

	switch event.Type(evt.Event) {
	case event.CreateUser:
		return CreateUserEventFromJSON(b)
	case event.CreateProject:
		return CreateProjectEventFromJSON(b)
	case event.UpdatePixel:
		return DrawPixelEventFromJSON(b)
	default:
		return nil, errors.New("invalid event type")
	}
}

func ToJSON(e event.Event) ([]byte, error) {
	switch e := e.(type) {
	case event.CreateUserEvent:
		return CreateUserEventToJSON(e)
	case event.CreateProjectEvent:
		return CreateProjectEventToJSON(e)
	case event.DrawPixelEvent:
		return DrawPixelEventToJSON(e)
	case event.UserCreatedEvent:
		return UserCreatedEventToJSON(e)
	case event.ProjectCreatedEvent:
		return ProjectCreatedEventToJSON(e)
	case event.PixelDrawnEvent:
		return PixelDrawnEventToJSON(e)
	default:
		return nil, errors.New(fmt.Sprint("no event marshaler", e))
	}
}
