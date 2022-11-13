package event

import "encoding/json"

const CreateProject Type = "create_project"

type CreateProjectEvent struct {
	Event
	Name string `json:"name"`
}

func CreateProjectEventFromJSON(b []byte) (CreateProjectEvent, error) {
	var e CreateProjectEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return CreateProjectEvent{}, err
	}

	return e, nil
}
