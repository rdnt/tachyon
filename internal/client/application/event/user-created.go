package event

import (
	"tachyon/pkg/uuid"
)

const (
	UserCreated Type = "user_created"
)

type UserCreatedEvent struct {
	UserId uuid.UUID
	Name   string
}

func (UserCreatedEvent) Type() Type {
	return UserCreated
}
