package event

import (
	"tachyon/pkg/uuid"
)

const (
	SessionCreated Type = "session_created"
	JoinedSession  Type = "joined_session"
	LeftSession    Type = "left_session"
)

type SessionCreatedEvent struct {
	ProjectId uuid.UUID
	SessionId uuid.UUID
	Name      string
	UserIds   []uuid.UUID
}

func (SessionCreatedEvent) Type() Type {
	return SessionCreated
}

//func (SessionCreatedEvent) AggregateType() AggregateType {
//	return Session
//}
//
//func (e SessionCreatedEvent) AggregateId() uuid.UUID {
//	return e.SessionId
//}

type JoinedSessionEvent struct {
	SessionId uuid.UUID
	UserId    uuid.UUID
}

func (JoinedSessionEvent) Type() Type {
	return JoinedSession
}

//func (JoinedSessionEvent) AggregateType() AggregateType {
//	return Session
//}
//
//func (e JoinedSessionEvent) AggregateId() uuid.UUID {
//	return e.SessionId
//}

type LeftSessionEvent struct {
	SessionId uuid.UUID
	UserId    uuid.UUID
}

func (LeftSessionEvent) Type() Type {
	return LeftSession
}

//func (LeftSessionEvent) AggregateType() AggregateType {
//	return Session
//}
//
//func (e LeftSessionEvent) AggregateId() uuid.UUID {
//	return e.SessionId
//}
