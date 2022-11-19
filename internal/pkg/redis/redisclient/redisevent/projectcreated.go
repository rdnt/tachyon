package redisevent

import (
	"encoding/json"

	"tachyon/internal/server/application/event"
	"tachyon/pkg/uuid"
)

type ProjectCreatedEvent struct {
	ProjectId string `json:"projectId"`
	OwnerId   string `json:"ownerId"`
	Name      string `json:"name"`
}

func ProjectCreatedEventToJSON(e event.ProjectCreatedEvent) ([]byte, error) {
	evt := ProjectCreatedEvent{
		ProjectId: e.ProjectId.String(),
		OwnerId:   e.OwnerId.String(),
		Name:      e.Name,
	}

	return json.Marshal(evt)
}

func ProjectCreatedEventFromJSON(b []byte) (event.ProjectCreatedEvent, error) {
	var e ProjectCreatedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return event.ProjectCreatedEvent{}, err
	}

	pid, err := uuid.Parse(e.ProjectId)
	if err != nil {
		return event.ProjectCreatedEvent{}, err
	}

	oid, err := uuid.Parse(e.OwnerId)
	if err != nil {
		return event.ProjectCreatedEvent{}, err
	}

	return event.ProjectCreatedEvent{
		ProjectId: pid,
		OwnerId:   oid,
		Name:      e.Name,
	}, nil
}
