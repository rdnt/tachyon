package websocketevent

import (
	"encoding/json"

	"tachyon/internal/client/application/event"
)

type ProjectCreatedEvent struct {
	Event     string `json:"event"`
	ProjectId string `json:"projectId"`
	OwnerId   string `json:"ownerId"`
	Name      string `json:"name"`
}

func ProjectCreatedEventToJSON(e event.ProjectCreatedEvent) ([]byte, error) {
	e2 := ProjectCreatedEvent{
		Event:     event.ProjectCreated.String(),
		ProjectId: e.ProjectId.String(),
		OwnerId:   e.OwnerId.String(),
		Name:      e.Name,
	}

	return json.Marshal(e2)
}
