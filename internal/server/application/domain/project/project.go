package project

import (
	"tachyon/internal/server/application/domain/project/path"
	"tachyon/pkg/uuid"
)

type Project struct {
	Id      uuid.UUID
	Name    string
	OwnerId uuid.UUID
	Paths   []path.Path
}

func New(id uuid.UUID, ownerId uuid.UUID, name string) Project {
	return Project{
		Id:      id,
		Name:    name,
		OwnerId: ownerId,
	}
}
