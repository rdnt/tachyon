// Package events provides shared definitions for server-created events that
// the client should know about
package event

import (
	"encoding/json"
	"errors"
)

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

	CreateProject  Type = "create-project"
	ProjectCreated Type = "project-created"

	CreatePath  Type = "create-path"
	PathCreated Type = "path-created"

	TracePth   Type = "trace-path"
	PathTraced Type = "path-traced"
)

// func MarshalJSON(e Event) ([]byte, error) {
//	switch e := e.(type) {
//	case event.UserCreatedEvent:
//		return UserCreatedEventToJSON(e)
//	case event.ProjectCreatedEvent:
//		return ProjectCreatedEventToJSON(e)
//	case event.SessionCreatedEvent:
//		return SessionCreatedEventToJSON(e)
//	case event.JoinedSessionEvent:
//		return JoinedSessionEventToJSON(e)
//	case event.LeftSessionEvent:
//		return LeftSessionEventToJSON(e)
//	case event.PixelDrawnEvent:
//		return PixelDrawnEventToJSON(e)
//	case event.PixelErasedEvent:
//		return PixelErasedEventToJSON(e)
//	default:
//		return nil, errors.New(fmt.Sprint("no event marshaler2", e))
//	}
// }

type jsonEvent struct {
	Event string `json:"event"`
}

func ToJSON(e Event) ([]byte, error) {
	b, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	var tmp map[string]any
	err = json.Unmarshal(b, &tmp)
	if err != nil {
		return nil, err
	}

	tmp["event"] = e.Type()

	return json.Marshal(tmp)
}

func FromJSON(b []byte) (Event, error) {
	var tmp jsonEvent
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return nil, err
	}

	switch Type(tmp.Event) {
	case Connected:
		return ConnectedEventFromJSON(b)
	case CreatePath:
		return CreatePathEventFromJSON(b)
	case CreateProject:
		return CreateProjectEventFromJSON(b)
	case CreateSession:
		return CreateSessionEventFromJSON(b)
	case JoinedSession:
		return JoinedSessionEventFromJSON(b)
	case LeftSession:
		return LeftSessionEventFromJSON(b)
	case PathCreated:
		return PathCreatedEventFromJSON(b)
	case ProjectCreated:
		return ProjectCreatedEventFromJSON(b)
	case SessionCreated:
		return SessionCreatedEventFromJSON(b)
	default:
		return nil, errors.New("invalid event type")
	}
}
