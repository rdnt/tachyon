package event

import "encoding/json"

type ProjectCreatedEvent struct {
	ProjectId string `json:"projectId"`
}

func (e ProjectCreatedEvent) Type() Type {
	return ProjectCreated
}

func (e ProjectCreatedEvent) AggregateType() AggregateType {
	return Project
}

func (e ProjectCreatedEvent) AggregateId() string {
	return e.ProjectId
}

func ProjectCreatedEventFromJSON(b []byte) (ProjectCreatedEvent, error) {
	var e ProjectCreatedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return ProjectCreatedEvent{}, err
	}

	return e, nil
}
