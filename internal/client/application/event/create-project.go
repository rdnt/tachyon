package event

const (
	CreateProject Type = "create_project"
)

type CreateProjectEvent struct {
	Name string
}

func (CreateProjectEvent) Type() Type {
	return CreateProject
}
