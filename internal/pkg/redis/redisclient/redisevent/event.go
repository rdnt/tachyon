package redisevent

import (
	"errors"
	"fmt"

	"tachyon/internal/server/application/event"
)

type Event struct {
	Type          string `json:"type"`
	AggregateType string `json:"aggregateType"`
	AggregateId   string `json:"aggregateId"`
}

func EventToJSON(e event.Event) ([]byte, error) {
	switch e := e.(type) {
	case event.UserCreatedEvent:
		return UserCreatedEventToJSON(e)
	case event.ProjectCreatedEvent:
		return ProjectCreatedEventToJSON(e)
	case event.SessionCreatedEvent:
		return SessionCreatedEventToJSON(e)
	case event.JoinedSessionEvent:
		return JoinedSessionEventToJSON(e)
	case event.LeftSessionEvent:
		return LeftSessionEventToJSON(e)
	case event.PathCreatedEvent:
		return PathCreatedEventToJSON(e)
	case event.PathTracedEvent:
		return PathTracedEventToJSON(e)
	default:
		return nil, errors.New(fmt.Sprint("no event marshaler2", e))
	}
}

func EventFromJSON(typ event.Type, b []byte) (event.Event, error) {
	switch typ {
	case event.UserCreated:
		return UserCreatedEventFromJSON(b)
	case event.ProjectCreated:
		return ProjectCreatedEventFromJSON(b)
	case event.SessionCreated:
		return SessionCreatedEventFromJSON(b)
	case event.JoinedSession:
		return JoinedSessionEventFromJSON(b)
	case event.LeftSession:
		return LeftSessionEventFromJSON(b)
	case event.PathCreated:
		return PathCreatedEventFromJSON(b)
	case event.PathTraced:
		return PathTracedEventFromJSON(b)
	default:
		return nil, errors.New("invalid event type")
	}
}
