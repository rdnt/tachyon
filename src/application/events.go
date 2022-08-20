package application

import (
	"tachyon2/pkg/logger"
	"tachyon2/src/application/event"
)

type Events struct {
	SessionJoined EventBroker[event.SessionJoinedEvent]
	SessionLeft   EventBroker[event.SessionLeftEvent]
	PathCreated   EventBroker[event.PathCreatedEvent]
	PathTraced    EventBroker[event.PathTracedEvent]
}

func newEvents(appId string) Events {
	log := logger.New("event-"+appId, logger.BlueFg)

	return Events{
		SessionJoined: newEventBroker[event.SessionJoinedEvent](log),
		SessionLeft:   newEventBroker[event.SessionLeftEvent](log),
		PathCreated:   newEventBroker[event.PathCreatedEvent](log),
		PathTraced:    newEventBroker[event.PathTracedEvent](log),
	}
}
