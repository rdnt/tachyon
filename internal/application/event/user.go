package event

import (
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/pkg/uuid"
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
