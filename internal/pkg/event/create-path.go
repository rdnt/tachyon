package event

type CreatePathEvent struct {
	ProjectId string  `json:"projectId"`
	Tool      string  `json:"tool"`
	Color     string  `json:"color"`
	Point     Vector2 `json:"point"`
}

func (e CreatePathEvent) Type() Type {
	return CreatePath
}
