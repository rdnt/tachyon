package event

type LeftSessionEvent struct {
	SessionId string `json:"sessionId"`
	UserId    string `json:"userId"`
}

func (e LeftSessionEvent) Type() Type {
	return LeftSession
}

func (e LeftSessionEvent) AggregateType() AggregateType {
	return Session
}

func (e LeftSessionEvent) AggregateId() string {
	return e.SessionId
}
