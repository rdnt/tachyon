package websocketevent

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/client/application/event"
)

type CreateProjectEvent struct {
	Name string `json:"name"`
}

func CreateProjectEventFromJSON(b []byte) (event.CreateProjectEvent, error) {
	var e CreateProjectEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return event.CreateProjectEvent{}, err
	}

	return event.CreateProjectEvent{Name: e.Name}, nil
}
