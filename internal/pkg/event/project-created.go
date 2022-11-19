package event

type ProjectCreatedEvent struct {
	ProjectId string `json:"projectId"`
}

func (e ProjectCreatedEvent) Type() Type {
	return ProjectCreated
}

func (e ProjectCreatedEvent) AggregateType() AggregateType {
	return Project
}

func (e ProjectCreatedEvent) AggregateId() string {
	return e.ProjectId
}
