package event

type CreateSessionEvent struct {
	Name string `json:"name"`
}

func (e CreateSessionEvent) Type() Type {
	return CreateSession
}
