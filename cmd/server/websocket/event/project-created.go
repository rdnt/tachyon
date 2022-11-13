package event

type ProjectCreatedEvent struct {
	Event
	ProjectId string `json:"projectId"`
	OwnerId   string `json:"ownerId"`
	Name      string `json:"name"`
}

//func ProjectCreatedEventToJSON(e ProjectCreatedEvent) ([]byte, error) {
//	e.Event.Event = "project_created"
//
//	return json.Marshal(e)
//}
