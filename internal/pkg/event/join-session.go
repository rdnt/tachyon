package event

type JoinSessionEvent struct {
	SessionId string `json:"sessionId"`
}

func (e JoinSessionEvent) Type() Type {
	return JoinSession
}
