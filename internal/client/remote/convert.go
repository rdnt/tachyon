package remote

import (
	"encoding/json"
	"errors"

	"tachyon/internal/pkg/event"
)

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

func EventFromJSON(b []byte) (Event, error) {
	var tmp jsonEvent
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return nil, err
	}

	switch event.Type(tmp.Event) {
	case event.Connected:
		return event.ConnectedEventFromJSON(b)
	case event.JoinedSession:
		return event.JoinedSessionEventFromJSON(b)
	case event.LeftSession:
		return event.LeftSessionEventFromJSON(b)
	case event.PathCreated:
		return event.PathCreatedEventFromJSON(b)
	case event.ProjectCreated:
		return event.ProjectCreatedEventFromJSON(b)
	case event.SessionCreated:
		return event.SessionCreatedEventFromJSON(b)
	default:
		return nil, errors.New("invalid event type")
	}
}
