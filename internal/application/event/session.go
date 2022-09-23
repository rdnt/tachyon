package event

import (
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/domain/user"
)

const (
	SessionCreated Type = "session_created"
	//JoinedSession  Type = "joined_session"
)

type SessionCreatedEvent struct {
	event

	ProjectId project.Id
	Id        session.Id
	Name      string
	OwnerId   user.Id
	UserIds   []user.Id
}

func NewSessionCreatedEvent(e SessionCreatedEvent) SessionCreatedEvent {
	e.typ = SessionCreated
	e.aggregateType = Session
	e.aggregateId = string(e.Id)

	return e
}

//type JoinedSessionEvent struct {
//	event
//
//	SessionId session.Id
//	UserId    user.Id
//}
//
//func NewJoinedSessionEvent(e JoinedSessionEvent) JoinedSessionEvent {
//	e.typ = JoinedSession
//	e.aggregateType = Session
//	e.aggregateId = string(e.SessionId)
//
//	return e
//}
