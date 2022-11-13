package event

import "encoding/json"

const DrawPixel Type = "draw_pixel"

type DrawPixelEvent struct {
	Event
	ProjectId string     `json:"projectId"`
	Color     string     `json:"color"`
	Coords    IntVector2 `json:"coords"`
}

func DrawPixelEventFromJSON(b []byte) (DrawPixelEvent, error) {
	var e DrawPixelEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return DrawPixelEvent{}, err
	}

	return e, nil
}
