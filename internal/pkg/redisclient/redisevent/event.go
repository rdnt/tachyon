package redisevent

import (
	"errors"

	"github.com/rdnt/tachyon/internal/application/event"
)

type Event struct {
	Type          string `json:"type"`
	AggregateType string `json:"aggregateType"`
	AggregateId   string `json:"aggregateId"`
}

func EventToJSON(e event.Event) ([]byte, error) {
	switch e := e.(type) {
	case event.PixelDrawnEvent:
		return PixelDrawnEventToJSON(e)
	case event.PixelErasedEvent:
		return PixelErasedEventToJSON(e)
	default:
		return nil, errors.New("no event marshaler")
	}
}

func EventFromJSON(typ event.Type, b []byte) (event.Event, error) {
	switch typ {
	case event.PixelDrawn:
		return PixelDrawnEventFromJSON(b)
	case event.PixelErased:
		return PixelErasedEventFromJSON(b)
	default:
		return nil, errors.New("invalid event type")
	}
}

//func redisEventFromJSON(e Event) (event.Event, error) {
//	if !slices.Contains(event.Types, event.Type(e.Type)) {
//		return event.Event{}, errors.New("invalid type")
//	}
//
//	if !slices.Contains(event.AggregateTypes, event.AggregateType(e.AggregateType)) {
//		return event.Event{}, errors.New("invalid type")
//	}
//
//	aggregateId, err := uuid.Parse(e.AggregateId)
//	if err != nil {
//		return event.Event{}, err
//	}
//
//	return event.New(
//		event.Type(e.Type),
//		event.AggregateType(e.AggregateType),
//		aggregateId,
//	), nil
//}
