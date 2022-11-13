package redisevent

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type JoinedSessionEvent struct {
	SessionId string `json:"sessionId"`
	UserId    string `json:"userId"`
}

func JoinedSessionEventToJSON(e event.JoinedSessionEvent) ([]byte, error) {
	evt := JoinedSessionEvent{
		SessionId: e.SessionId.String(),
		UserId:    e.UserId.String(),
	}

	return json.Marshal(evt)
}

func JoinedSessionEventFromJSON(b []byte) (event.JoinedSessionEvent, error) {
	var e JoinedSessionEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return event.JoinedSessionEvent{}, err
	}

	sid, err := uuid.Parse(e.SessionId)
	if err != nil {
		return event.JoinedSessionEvent{}, err
	}

	uid, err := uuid.Parse(e.UserId)
	if err != nil {
		return event.JoinedSessionEvent{}, err
	}

	return event.JoinedSessionEvent{
		SessionId: sid,
		UserId:    uid,
	}, nil
}
