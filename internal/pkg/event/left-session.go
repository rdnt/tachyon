package event

import "encoding/json"

type LeftSessionEvent struct {
	SessionId string `json:"sessionId"`
	UserId    string `json:"userId"`
}

func (e LeftSessionEvent) Type() Type {
	return LeftSession
}

func (e LeftSessionEvent) AggregateType() AggregateType {
	return Session
}

func (e LeftSessionEvent) AggregateId() string {
	return e.SessionId
}

func LeftSessionEventFromJSON(b []byte) (LeftSessionEvent, error) {
	var e LeftSessionEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return LeftSessionEvent{}, err
	}

	return e, nil
}
