package redisevent

import (
	"encoding/json"

	"tachyon/internal/server/application/event"
	"tachyon/pkg/uuid"
)

type UserCreatedEvent struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

func UserCreatedEventToJSON(e event.UserCreatedEvent) ([]byte, error) {
	evt := UserCreatedEvent{
		UserId: e.UserId.String(),
		Name:   e.Name,
	}

	return json.Marshal(evt)
}

func UserCreatedEventFromJSON(b []byte) (event.UserCreatedEvent, error) {
	var evt UserCreatedEvent
	err := json.Unmarshal(b, &evt)
	if err != nil {
		return event.UserCreatedEvent{}, err
	}

	uid, err := uuid.Parse(evt.UserId)
	if err != nil {
		return event.UserCreatedEvent{}, err
	}

	return event.UserCreatedEvent{
		UserId: uid,
		Name:   evt.Name,
	}, nil
}
