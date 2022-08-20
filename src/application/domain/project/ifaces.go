package project

type Repository interface {
	CreateProject(Project) (Project, error)
	Project(Id) (Project, error)
	Projects() (map[Id]Project, error)
	DeleteProject(Id) error
	UpdateProject(Project) (Project, error)
}
