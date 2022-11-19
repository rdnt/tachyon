package event

import (
	"errors"

	"github.com/rdnt/tachyon/pkg/uuid"
	"golang.org/x/exp/slices"
)

type Event interface {
	Type() Type
	AggregateType() AggregateType
	AggregateId() uuid.UUID
}

type AggregateType string

func (t AggregateType) String() string {
	return string(t)
}

const (
	User    AggregateType = "user"
	Project AggregateType = "project"
	Session AggregateType = "session"
)

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
