package event

import (
	"encoding/json"
	"tachyon/pkg/uuid"
)

type CreateProjectEvent struct {
	Name string `json:"name"`
}

func (e CreateProjectEvent) Type() Type {
	return CreateProject
}

func (e CreateProjectEvent) AggregateType() AggregateType {
	return Project
}

func (e CreateProjectEvent) AggregateId() string {
	return uuid.Nil.String()
}

func CreateProjectEventFromJSON(b []byte) (CreateProjectEvent, error) {
	var e CreateProjectEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return CreateProjectEvent{}, err
	}

	return e, nil
}
