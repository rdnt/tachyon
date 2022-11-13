package event

type PixelDrawnEvent struct {
	Event
	UserId    string     `json:"userId"`
	ProjectId string     `json:"projectId"`
	Color     string     `json:"color"`
	Coords    IntVector2 `json:"coords"`
}

//func PixelDrawnEventToJSON(e PixelDrawnEvent) ([]byte, error) {
//	e.Event.Event = "pixel_drawn"
//
//	return json.Marshal(e)
//}
