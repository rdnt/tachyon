package event

import "encoding/json"

type ConnectedEvent struct {
	UserId string `json:"userId"`
}

func (e ConnectedEvent) Type() Type {
	return Connected
}

func (e ConnectedEvent) AggregateType() AggregateType {
	return User
}

func (e ConnectedEvent) AggregateId() string {
	return e.UserId
}

func ConnectedEventFromJSON(b []byte) (ConnectedEvent, error) {
	var e ConnectedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return ConnectedEvent{}, err
	}

	return e, nil
}
