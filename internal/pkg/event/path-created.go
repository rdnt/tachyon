package event

import "encoding/json"

type PathCreatedEvent struct {
	PathId    string  `json:"pathId"`
	UserId    string  `json:"userId"`
	ProjectId string  `json:"projectId"`
	Tool      string  `json:"tool"`
	Color     string  `json:"color"`
	Point     Vector2 `json:"point"`
}

func (e PathCreatedEvent) Type() Type {
	return PathCreated
}

func (e PathCreatedEvent) AggregateType() AggregateType {
	return Project
}

func (e PathCreatedEvent) AggregateId() string {
	return e.ProjectId
}

func PathCreatedEventFromJSON(b []byte) (PathCreatedEvent, error) {
	var e PathCreatedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return PathCreatedEvent{}, err
	}

	return e, nil
}
