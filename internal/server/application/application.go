package application

import (
	"github.com/rdnt/tachyon/internal/server/application/domain/session"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type EventStore[E any] interface {
	Publish(event E) (err error)
	Subscribe(handler func(E)) (dispose func(), err error)
	Events() (events []E, err error)
}

type EventBus[E any] interface {
	Publish(event E) (err error)
	Subscribe(handler func(E)) (dispose func(), err error)
	Events() (events []E, err error)
}

type Commands interface {
	CreateSession(userId uuid.UUID, projectId uuid.UUID, sessionName string) error
	JoinSession(userId uuid.UUID, sessionId uuid.UUID) error
}

type Queries interface {
	Session(id uuid.UUID) (session.Session, error)
}
