package application

import (
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/domain/user"
)

type Commands interface {
	CreateSession(userId user.Id, projectId project.Id, sessionName string) error
	JoinSession(userId user.Id, sessionId session.Id) error
}

type Queries interface {
	Session(id session.Id) (session.Session, error)
}

type App struct {
	Commands
	Queries
}

func New(cmds Commands, qs Queries) *App {
	return &App{
		Commands: cmds,
		Queries:  qs,
	}
}
