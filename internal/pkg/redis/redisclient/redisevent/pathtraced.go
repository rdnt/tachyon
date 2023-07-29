package redisevent

import (
	"encoding/json"

	"tachyon/internal/server/application/domain/project/path"
	"tachyon/internal/server/application/event"
	"tachyon/pkg/uuid"
)

type PathTracedEvent struct {
	PathId    string  `json:"pathId"`
	UserId    string  `json:"userId"`
	ProjectId string  `json:"projectId"`
	Point     Vector2 `json:"coords"`
}

func PathTracedEventToJSON(e event.PathTracedEvent) ([]byte, error) {
	evt := PathTracedEvent{
		PathId:    e.PathId.String(),
		UserId:    e.UserId.String(),
		ProjectId: e.ProjectId.String(),
		Point: Vector2{
			X: e.Point.X,
			Y: e.Point.Y,
		},
	}

	return json.Marshal(evt)
}

func PathTracedEventFromJSON(b []byte) (event.PathTracedEvent, error) {
	var evt PathTracedEvent
	err := json.Unmarshal(b, &evt)
	if err != nil {
		return event.PathTracedEvent{}, err
	}

	pathId, err := uuid.Parse(evt.PathId)
	if err != nil {
		return event.PathTracedEvent{}, err
	}

	uid, err := uuid.Parse(evt.UserId)
	if err != nil {
		return event.PathTracedEvent{}, err
	}

	pid, err := uuid.Parse(evt.ProjectId)
	if err != nil {
		return event.PathTracedEvent{}, err
	}

	return event.PathTracedEvent{
		PathId:    pathId,
		UserId:    uid,
		ProjectId: pid,
		Point: path.Vector2{
			X: evt.Point.X,
			Y: evt.Point.Y,
		},
	}, nil
}
