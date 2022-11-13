package event

type UserCreatedEvent struct {
	Event
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

//func UserCreatedEventToJSON(e UserCreatedEvent) ([]byte, error) {
//	e.Event.Event = "project_created"
//	return json.Marshal(e)
//}
