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

type SessionJoinedEvent struct {
	SessionId session.Id `json:"sessionId"`
	UserId    user.Id    `json:"userId"`
}

func (SessionJoinedEvent) Type() Type {
	return SessionJoined
}

type SessionLeftEvent struct {
	SessionId session.Id `json:"sessionId"`
	UserId    user.Id    `json:"userId"`
}

func (SessionLeftEvent) Type() Type {
	return SessionLeft
}

type PathCreatedEvent struct {
	SessionId session.Id `json:"sessionId"`
	UserId    user.Id    `json:"userId"`
	Path      path.Path  `json:"path"`
}

//
// type Path struct {
// 	Id     string  `json:"id"`
// 	Tool   string  `json:"string"`
// 	Color  string  `json:"string"`
// 	Points []Point `json:"points"`
// }
//
// type Point struct {
// 	X float64 `json:"x"`
// 	Y float64 `json:"y"`
// }

func (PathCreatedEvent) Type() Type {
	return PathCreated
}

type PathTracedEvent struct {
	SessionId session.Id `json:"sessionId"`
	UserId    user.Id    `json:"userId"`
}

func (PathTracedEvent) Type() Type {
	return PathTraced
}
