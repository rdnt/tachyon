package websocketevent

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/client/application/event"
)

type UserCreatedEvent struct {
	Event  string `json:"event"`
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

func UserCreatedEventToJSON(e event.UserCreatedEvent) ([]byte, error) {
	e2 := UserCreatedEvent{
		Event:  event.UserCreated.String(),
		UserId: e.UserId.String(),
		Name:   e.Name,
	}

	return json.Marshal(e2)
}
