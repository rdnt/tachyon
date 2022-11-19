package event

type UpdatePixelEvent struct {
	ProjectId string  `json:"projectId"`
	Color     string  `json:"color"`
	Coords    Vector2 `json:"coords"`
}

func (e UpdatePixelEvent) Type() Type {
	return UpdatePixel
}
