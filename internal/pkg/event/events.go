// Package events provides shared definitions for server-created events that
// the client should know about
package event

type Event interface {
	Type() Type
	// AggregateType() AggregateType
	// AggregateId() string
}

type AggregateType string

const (
	User    AggregateType = "user"
	Project AggregateType = "project"
	Session AggregateType = "session"
)

type Type string

const (
	Connected Type = "connected"

	CreateSession  Type = "create-session"
	SessionCreated Type = "session-created"

	JoinSession   Type = "join-session"
	JoinedSession Type = "joined-session"

	LeaveSession Type = "leave-session"
	LeftSession  Type = "left-session"

	ProjectCreated Type = "project-created"

	UpdatePixel  Type = "pixel-updated"
	PixelUpdated Type = "pixel-updated"
)
