package event

import "github.com/google/uuid"

type Event interface {
	Type() Type
	AggregateType() AggregateType
	AggregateId() uuid.UUID
}

type event struct {
	typ           Type
	aggregateType AggregateType
	aggregateId   uuid.UUID
}

func (e event) Type() Type {
	return e.typ
}

func (e event) AggregateType() AggregateType {
	return e.aggregateType
}

func (e event) AggregateId() uuid.UUID {
	return e.aggregateId
}

type AggregateType string

const (
	Session AggregateType = "session"
	User    AggregateType = "user"
	Project AggregateType = "project"
)

type Type string
