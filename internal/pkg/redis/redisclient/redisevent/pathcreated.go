package redisevent

import (
	"encoding/json"

	"tachyon/internal/server/application/domain/project/path"
	"tachyon/internal/server/application/event"
	"tachyon/pkg/uuid"
)

type PathCreatedEvent struct {
	PathId    string  `json:"pathId"`
	UserId    string  `json:"userId"`
	ProjectId string  `json:"projectId"`
	Tool      string  `json:"tool"`
	Color     string  `json:"color"`
	Point     Vector2 `json:"coords"`
}

func PathCreatedEventToJSON(e event.PathCreatedEvent) ([]byte, error) {
	evt := PathCreatedEvent{
		PathId:    e.PathId.String(),
		UserId:    e.UserId.String(),
		ProjectId: e.ProjectId.String(),
		Tool:      e.Tool.String(),
		Color:     e.Color.String(),
		Point: Vector2{
			X: e.Point.X,
			Y: e.Point.Y,
		},
	}

	return json.Marshal(evt)
}

func PathCreatedEventFromJSON(b []byte) (event.PathCreatedEvent, error) {
	var evt PathCreatedEvent
	err := json.Unmarshal(b, &evt)
	if err != nil {
		return event.PathCreatedEvent{}, err
	}

	pathId, err := uuid.Parse(evt.PathId)
	if err != nil {
		return event.PathCreatedEvent{}, err
	}

	uid, err := uuid.Parse(evt.UserId)
	if err != nil {
		return event.PathCreatedEvent{}, err
	}

	pid, err := uuid.Parse(evt.ProjectId)
	if err != nil {
		return event.PathCreatedEvent{}, err
	}

	clr, err := path.ColorFromString(evt.Color)
	if err != nil {
		return event.PathCreatedEvent{}, err
	}

	return event.PathCreatedEvent{
		PathId:    pathId,
		UserId:    uid,
		ProjectId: pid,
		Tool:      path.Tool(evt.Tool),
		Color:     clr,
		Point: path.Vector2{
			X: evt.Point.X,
			Y: evt.Point.Y,
		},
	}, nil
}
