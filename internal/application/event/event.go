package event

type Event interface {
	Type() Type
	AggregateType() AggregateType
	AggregateId() string
}

type event struct {
	typ           Type
	aggregateType AggregateType
	aggregateId   string
}

func (e event) Type() Type {
	return e.typ
}

func (e event) AggregateType() AggregateType {
	return e.aggregateType
}

func (e event) AggregateId() string {
	return e.aggregateId
}

type AggregateType string

const (
	Session AggregateType = "session"
	User    AggregateType = "user"
	//Project AggregateType = "project"
)

type Type string
