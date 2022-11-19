package event

type ConnectedEvent struct {
	UserId string `json:"userId"`
}

func (e ConnectedEvent) Type() Type {
	return Connected
}

func (e ConnectedEvent) AggregateType() AggregateType {
	return User
}

func (e ConnectedEvent) AggregateId() string {
	return e.UserId
}
