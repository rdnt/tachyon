// Package events provides shared definitions for server-created events that
// the client should know about
package event

type Event interface {
	Type() Type
	AggregateType() AggregateType
	AggregateId() string
}

type AggregateType string

const (
	User    AggregateType = "user"
	Project AggregateType = "project"
	Session AggregateType = "session"
)

type Type string

const (
	UserCreated    Type = "user-created"
	ProjectCreated Type = "project-created"
	SessionCreated Type = "session-created"
	JoinedSession  Type = "joined-session"
	LeftSession    Type = "left-session"
	PixelUpdated   Type = "pixel-updated"
)
