package event

import "github.com/google/uuid"

type EventRWIface interface {
	Type() Type
	SetType(typ Type)
	AggregateType() AggregateType
	SetAggregateType(typ AggregateType)
	AggregateId() uuid.UUID
	SetAggregateId(id uuid.UUID)
}

type EventIface interface{}

type Event struct {
	Type          Type
	AggregateType AggregateType
	AggregateId   uuid.UUID
}

type AggregateType string

const (
	Session AggregateType = "session"
	User    AggregateType = "user"
	Project AggregateType = "project"
)

type Type string
