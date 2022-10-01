package project

import (
	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/domain/user"
)

type Id uuid.UUID

type Project struct {
	Id      Id
	Name    string
	OwnerId user.Id
}

func New(id Id, ownerId user.Id, name string) Project {
	return Project{
		Id:      id,
		Name:    name,
		OwnerId: ownerId,
	}
}
