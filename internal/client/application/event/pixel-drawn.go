package event

type PixelUpdatedEvent struct {
	UserId    string  `json:"userId"`
	ProjectId string  `json:"projectId"`
	Color     string  `json:"color"`
	Coords    Vector2 `json:"coords"`
}

func (PixelUpdatedEvent) Type() Type {
	return PixelUpdated
}
