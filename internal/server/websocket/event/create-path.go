package event

import "encoding/json"

const CreatePath Type = "create-path"

type CreatePathEvent struct {
	Event
	PathId    string  `json:"pathId"`
	ProjectId string  `json:"projectId"`
	Tool      string  `json:"tool"`
	Color     string  `json:"color"`
	Point     Vector2 `json:"point"`
}

func CreatePathEventFromJSON(b []byte) (CreatePathEvent, error) {
	var e CreatePathEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return CreatePathEvent{}, err
	}

	return e, nil
}
