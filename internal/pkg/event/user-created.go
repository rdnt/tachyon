package event

type UserCreatedEvent struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

func (e UserCreatedEvent) Type() Type {
	return UserCreated
}

func (e UserCreatedEvent) AggregateType() AggregateType {
	return User
}

func (e UserCreatedEvent) AggregateId() string {
	return e.UserId
}
