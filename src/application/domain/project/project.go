package project

import (
	"tachyon2/pkg/uuid"
	"tachyon2/src/application/domain/project/path"
	"tachyon2/src/application/domain/user"
)

type Id string

type Project struct {
	Id                   Id
	Name                 string
	OwnerId              user.Id
	Paths                []path.Path
	UserHistory          map[user.Id][]path.Id
	UserHistoryIndicator map[user.Id]int
}

func New(name string, ownerId user.Id) Project {
	id := uuid.New()

	p := Project{
		Id:                   Id(id),
		Name:                 name,
		OwnerId:              ownerId,
		UserHistory:          map[user.Id][]path.Id{},
		UserHistoryIndicator: map[user.Id]int{},
	}

	// TODO: dont hardcode
	p.Id = "my-project"
	p.Name = "my-project"

	return p
}
