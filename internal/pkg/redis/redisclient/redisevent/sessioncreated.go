package redisevent

import (
	"encoding/json"

	"github.com/rdnt/tachyon/internal/server/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type SessionCreatedEvent struct {
	ProjectId string   `json:"projectId"`
	SessionId string   `json:"sessionId"`
	Name      string   `json:"name"`
	UserIds   []string `json:"userIds"`
}

func SessionCreatedEventToJSON(e event.SessionCreatedEvent) ([]byte, error) {
	ids := make([]string, len(e.UserIds))
	for i, id := range e.UserIds {
		ids[i] = id.String()
	}

	evt := SessionCreatedEvent{
		ProjectId: e.ProjectId.String(),
		SessionId: e.SessionId.String(),
		Name:      e.Name,
		UserIds:   ids,
	}

	return json.Marshal(evt)
}

func SessionCreatedEventFromJSON(b []byte) (event.SessionCreatedEvent, error) {
	var e SessionCreatedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return event.SessionCreatedEvent{}, err
	}

	pid, err := uuid.Parse(e.ProjectId)
	if err != nil {
		return event.SessionCreatedEvent{}, err
	}

	sid, err := uuid.Parse(e.SessionId)
	if err != nil {
		return event.SessionCreatedEvent{}, err
	}

	uids := make([]uuid.UUID, len(e.UserIds))
	for i, id := range e.UserIds {
		uid, err := uuid.Parse(id)
		if err != nil {
			return event.SessionCreatedEvent{}, err
		}

		uids[i] = uid
	}

	return event.SessionCreatedEvent{
		ProjectId: pid,
		SessionId: sid,
		Name:      e.Name,
		UserIds:   uids,
	}, nil
}
