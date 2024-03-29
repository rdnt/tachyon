package event

import "encoding/json"

type SessionCreatedEvent struct {
	ProjectId string `json:"projectId"`
	SessionId string `json:"sessionId"`
	Name      string `json:"name"`
}

func (e SessionCreatedEvent) Type() Type {
	return SessionCreated
}

func (e SessionCreatedEvent) AggregateType() AggregateType {
	return Session
}

func (e SessionCreatedEvent) AggregateId() string {
	return e.SessionId
}

func SessionCreatedEventFromJSON(b []byte) (SessionCreatedEvent, error) {
	var e SessionCreatedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return SessionCreatedEvent{}, err
	}

	return e, nil
}
