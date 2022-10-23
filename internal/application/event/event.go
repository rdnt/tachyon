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

type EventIface interface {
	Type() Type
	AggregateType() AggregateType
	AggregateId() uuid.UUID
}

type Event struct {
	typ           Type
	aggregateType AggregateType
	aggregateId   uuid.UUID
}

func (e Event) Type() Type {
	return e.typ
}

func (e *Event) SetType(typ Type) {
	e.typ = typ
}

func (e Event) AggregateType() AggregateType {
	return e.aggregateType
}

func (e *Event) SetAggregateType(aggregateType AggregateType) {
	e.aggregateType = aggregateType
}

func (e Event) AggregateId() uuid.UUID {
	return e.aggregateId
}

func (e *Event) SetAggregateId(aggregateId uuid.UUID) {
	e.aggregateId = aggregateId
}

type AggregateType string

const (
	Session AggregateType = "session"
	User    AggregateType = "user"
	Project AggregateType = "project"
)

type Type string
