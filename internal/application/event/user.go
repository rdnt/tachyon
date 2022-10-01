package event

import (
	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/domain/user"
)

const (
	UserCreated Type = "user_created"
)

type UserCreatedEvent struct {
	event

	Id   user.Id
	Name string
}

func NewUserCreatedEvent(e UserCreatedEvent) UserCreatedEvent {
	e.typ = UserCreated
	e.aggregateType = User
	e.aggregateId = uuid.UUID(e.Id)

	return e
}
