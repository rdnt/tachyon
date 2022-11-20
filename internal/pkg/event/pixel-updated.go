package event

import "encoding/json"

type PixelUpdatedEvent struct {
	UserId    string  `json:"userId"`
	ProjectId string  `json:"projectId"`
	Color     string  `json:"color"`
	Coords    Vector2 `json:"coords"`
}

func (e PixelUpdatedEvent) Type() Type {
	return PixelUpdated
}

func (e PixelUpdatedEvent) AggregateType() AggregateType {
	return Project
}

func (e PixelUpdatedEvent) AggregateId() string {
	return e.ProjectId
}

func PixelUpdatedEventFromJSON(b []byte) (PixelUpdatedEvent, error) {
	var e PixelUpdatedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return PixelUpdatedEvent{}, err
	}

	return e, nil
}
