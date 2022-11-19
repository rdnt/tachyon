package event

type ProjectCreatedEvent struct {
	ProjectId string `json:"projectId"`
	OwnerId   string `json:"ownerId"`
	Name      string `json:"name"`
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
