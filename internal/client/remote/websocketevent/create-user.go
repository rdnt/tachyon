package websocketevent

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/client/application/event"
)

type CreateUserEvent struct {
	Name string `json:"name"`
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
