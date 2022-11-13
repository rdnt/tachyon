package websocketevent

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/client/application/event"
)

type CreateUserEvent struct {
	Event string `json:"event"`
	Name  string `json:"name"`
}

func CreateUserEventFromJSON(b []byte) (event.CreateUserEvent, error) {
	var e CreateUserEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return event.CreateUserEvent{}, err
	}

	return event.CreateUserEvent{
		Name: e.Name,
	}, nil
}

func CreateUserEventToJSON(e event.CreateUserEvent) ([]byte, error) {
	e2 := CreateUserEvent{
		Event: e.Type().String(),
		Name:  e.Name,
	}

	return json.Marshal(e2)
}
