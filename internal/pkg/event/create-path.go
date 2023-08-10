package event

import "encoding/json"

type CreatePathEvent struct {
	ProjectId string  `json:"projectId"`
	Tool      string  `json:"tool"`
	Color     string  `json:"color"`
	Point     Vector2 `json:"point"`
}

func (e CreatePathEvent) Type() Type {
	return CreatePath
}

func CreatePathEventFromJSON(b []byte) (CreatePathEvent, error) {
	var e CreatePathEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return CreatePathEvent{}, err
	}

	return e, nil
}
