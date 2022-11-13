package websocketevent

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/client/application/event"
)

type CreateProjectEvent struct {
	Event string `json:"event"`
	Name  string `json:"name"`
}

func CreateProjectEventFromJSON(b []byte) (event.CreateProjectEvent, error) {
	var e CreateProjectEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return event.CreateProjectEvent{}, err
	}

	return event.CreateProjectEvent{Name: e.Name}, nil
}

func CreateProjectEventToJSON(e event.CreateProjectEvent) ([]byte, error) {
	e2 := CreateProjectEvent{
		Event: e.Type().String(),
		Name:  e.Name,
	}

	return json.Marshal(e2)
}
