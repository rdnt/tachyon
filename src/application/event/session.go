package event

import (
	"tachyon2/src/application/domain/project/path"
	"tachyon2/src/application/domain/session"
	"tachyon2/src/application/domain/user"
)

type Type string

const (
	SessionJoined Type = "session.joined"
	SessionLeft   Type = "session.left"
	PathCreated   Type = "path.created"
	PathTraced    Type = "path.traced"
)

type Event struct {
	AggregateId string
}

type SessionJoinedEvent struct {
	Event
	SessionId session.Id
	UserId    user.Id
}

type SessionLeftEvent struct {
	Event
	SessionId session.Id
	UserId    user.Id
}

type PathCreatedEvent struct {
	Event
	SessionId session.Id
	UserId    user.Id
	Path      path.Path
}

type PathTracedEvent struct {
	Event
	SessionId session.Id
	UserId    user.Id
}
