package event

type SessionCreatedEvent struct {
	ProjectId string   `json:"projectId"`
	SessionId string   `json:"sessionId"`
	Name      string   `json:"name"`
	UserIds   []string `json:"userIds"`
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
