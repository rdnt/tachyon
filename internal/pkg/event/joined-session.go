package event

type JoinedSessionEvent struct {
	SessionId string `json:"sessionId"`
	UserId    string `json:"userId"`
}

func (e JoinedSessionEvent) Type() Type {
	return JoinedSession
}

func (e JoinedSessionEvent) AggregateType() AggregateType {
	return Session
}

func (e JoinedSessionEvent) AggregateId() string {
	return e.SessionId
}
