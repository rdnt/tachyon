package application

import (
	"tachyon2/pkg/broker"
	"tachyon2/pkg/logger"
)

func newEventBroker[E any](log *logger.Logger) EventBroker[E] {
	return EventBroker[E]{
		broker: broker.New[E](),
		logger: log,
	}
}

type EventBroker[E any] struct {
	broker *broker.Broker[E]
	logger *logger.Logger
}

func (b EventBroker[E]) Subscribe(h func(e E)) (dispose func()) {
	return b.broker.Subscribe(h)
}

func (b EventBroker[E]) publish(e E) {
	b.logger.Printf("event published: %#v\n", e)

	b.broker.Publish(e)
}
