package event

import (
	"encoding/json"

	"tachyon/pkg/uuid"
)

type CreateSessionEvent struct {
	Name      string `json:"name"`
	ProjectId string `json:"projectId"`
}

func (e CreateSessionEvent) Type() Type {
	return CreateSession
}

func (e CreateSessionEvent) AggregateType() AggregateType {
	return Session
}

func (e CreateSessionEvent) AggregateId() string {
	return uuid.Nil.String()
}

func CreateSessionEventFromJSON(b []byte) (CreateSessionEvent, error) {
	var e CreateSessionEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return CreateSessionEvent{}, err
	}

	return e, nil
}
