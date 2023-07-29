package event

import (
	"errors"
)

type Type string

type Event struct {
	Event string `json:"event"`
}

func FromJSON(typ Type, b []byte) (any, error) {
	switch typ {
	case CreateUser:
		return CreateUserEventFromJSON(b)
	case CreateProject:
		return CreateProjectEventFromJSON(b)
	case CreatePath:
		return CreatePathEventFromJSON(b)
	default:
		return nil, errors.New("invalid event type")
	}
}

//func ToJSON(e any) ([]byte, error) {
//	return json.Marshal(e)
//}
