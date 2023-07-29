package event

import "encoding/json"

const CreateUser Type = "create-user"

type CreateUserEvent struct {
	Event
	Name string `json:"name"`
}

func CreateUserEventFromJSON(b []byte) (CreateUserEvent, error) {
	var e CreateUserEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return CreateUserEvent{}, err
	}

	return e, nil
}
