package event

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type LeftSessionEvent struct {
	SessionId string `json:"sessionId"`
	UserId    string `json:"userId"`
}

func LeftSessionEventToJSON(e event.LeftSessionEvent) ([]byte, error) {
	evt := LeftSessionEvent{
		SessionId: e.SessionId.String(),
		UserId:    e.UserId.String(),
	}

	return json.Marshal(evt)
}

func LeftSessionEventFromJSON(b []byte) (event.LeftSessionEvent, error) {
	var e LeftSessionEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return event.LeftSessionEvent{}, err
	}

	sid, err := uuid.Parse(e.SessionId)
	if err != nil {
		return event.LeftSessionEvent{}, err
	}

	uid, err := uuid.Parse(e.UserId)
	if err != nil {
		return event.LeftSessionEvent{}, err
	}

	return event.LeftSessionEvent{
		SessionId: sid,
		UserId:    uid,
	}, nil
}
