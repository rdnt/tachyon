package session

import (
	"fmt"

	"tachyon2/pkg/uuid"
	"tachyon2/src/application/domain/project"
	"tachyon2/src/application/domain/user"
)

type Id string

type Session struct {
	Id        Id
	Name      string
	ProjectId project.Id
	OwnerId   user.Id
	UserIds   []user.Id
}

func (s Session) String() string {
	return fmt.Sprintf("--- session{id: %s, name: %s, projectId: %s, ownerId: %s, memberIds: %s}", s.Id, s.Name, s.ProjectId, s.OwnerId, s.UserIds)
}

func New(name string, projectId project.Id, ownerId user.Id, memberIds ...user.Id) Session {
	id := uuid.New()

	return Session{
		Id:        Id(id),
		Name:      name,
		ProjectId: projectId,
		OwnerId:   ownerId,
		UserIds:   memberIds,
	}
}
