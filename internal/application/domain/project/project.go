package project

import "github.com/rdnt/tachyon/internal/application/domain/project/path"

type Id string

type Project struct {
	Id    Id
	Paths []path.Id
}

func New(id Id, paths []path.Id) Project {
	return Project{
		Id:    id,
		Paths: paths,
	}
}
