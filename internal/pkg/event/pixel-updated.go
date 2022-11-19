package event

type PixelUpdatedEvent struct {
	UserId    string  `json:"userId"`
	ProjectId string  `json:"projectId"`
	Color     string  `json:"color"`
	Coords    Vector2 `json:"coords"`
}

func (e PixelUpdatedEvent) Type() Type {
	return PixelUpdated
}

func (e PixelUpdatedEvent) AggregateType() AggregateType {
	return Project
}

func (e PixelUpdatedEvent) AggregateId() string {
	return e.ProjectId
}
