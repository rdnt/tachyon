package event

import (
	"tachyon/pkg/uuid"
)

const (
	UserCreated Type = "user-created"
)

type UserCreatedEvent struct {
	UserId uuid.UUID
	Name   string
}

func (UserCreatedEvent) Type() Type {
	return UserCreated
}

func (UserCreatedEvent) AggregateType() AggregateType {
	return User
}

func (e UserCreatedEvent) AggregateId() uuid.UUID {
	return e.UserId
}
