package event

import (
	"errors"

	"tachyon/pkg/uuid"
	"golang.org/x/exp/slices"
)

//type Event interface {
//	Type() Type
//	SetType(typ Type)
//	AggregateType() AggregateType
//	SetAggregateType(typ AggregateType)
//	AggregateId() uuid.UUID
//	SetAggregateId(id uuid.UUID)
//}

type Event interface {
	Type() Type
	AggregateType() AggregateType
	AggregateId() uuid.UUID
}

//type Event struct {
//	Type          Type
//	AggregateType AggregateType
//	AggregateId   uuid.UUID
//}

//func (e Event) Type() Type {
//	return e.typ
//}
//
//func (e Event) AggregateType() AggregateType {
//	return e.aggregateType
//}
//
//func (e Event) AggregateId() uuid.UUID {
//	return e.aggregateId
//}

//func New(typ Type, aggr AggregateType, agrId uuid.UUID) Event {
//	return Event{
//		Type:          typ,
//		AggregateType: aggr,
//		AggregateId:   agrId,
//	}
//}

type AggregateType string

func (t AggregateType) String() string {
	return string(t)
}

const (
	User    AggregateType = "user"
	Project AggregateType = "project"
	Session AggregateType = "session"
)

//var AggregateTypes = []AggregateType{
//	User,
//	Project,
//	Session,
//}

type Type string

func (t Type) String() string {
	return string(t)
}

func TypeFromString(s string) (Type, error) {
	if !slices.Contains(Types, Type(s)) {
		return "", errors.New("invalid event type")
	}

	return Type(s), nil
}

var Types = []Type{
	UserCreated,
	ProjectCreated,
	SessionCreated,
	JoinedSession,
	LeftSession,
	PixelDrawn,
	PixelErased,
}
