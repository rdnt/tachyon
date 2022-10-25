package event

import "github.com/rdnt/tachyon/pkg/uuid"

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
	return e.ag
}

type AggregateType string

const (
	Session AggregateType = "session"
	User    AggregateType = "user"
	Project AggregateType = "project"
)

type Type string
