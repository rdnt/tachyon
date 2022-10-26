package redisevent

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type UserCreatedEvent struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

func UserCreatedEventToJSON(e event.UserCreatedEvent) ([]byte, error) {
	evt := UserCreatedEvent{
		UserId: uuid.UUID(e.UserId).String(),
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
		UserId: uuid.UUID(uid),
		Name:   evt.Name,
	}, nil
}
