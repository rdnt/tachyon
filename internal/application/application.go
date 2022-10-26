package application

import (
	"github.com/rdnt/tachyon/internal/application/domain/session"
)

type Commands interface {
	CreateSession(userId uuid.UUID, projectId uuid.UUID, sessionName string) error
	JoinSession(userId uuid.UUID, sessionId uuid.UUID) error
}

type Queries interface {
	Session(id uuid.UUID) (session.Session, error)
}
