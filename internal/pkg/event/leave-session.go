package event

type LeaveSessionEvent struct{}

func (e LeaveSessionEvent) Type() Type {
	return LeaveSession
}
