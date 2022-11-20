package event

import "encoding/json"

type JoinedSessionEvent struct {
	SessionId string `json:"sessionId"`
	UserId    string `json:"userId"`
}

func (e JoinedSessionEvent) Type() Type {
	return JoinedSession
}

func (e JoinedSessionEvent) AggregateType() AggregateType {
	return Session
}

func (e JoinedSessionEvent) AggregateId() string {
	return e.SessionId
}

func JoinedSessionEventFromJSON(b []byte) (JoinedSessionEvent, error) {
	var e JoinedSessionEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return JoinedSessionEvent{}, err
	}

	return e, nil
}
