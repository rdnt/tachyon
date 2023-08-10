package application

import (
	"errors"
	"time"

	"tachyon/internal/client/application/domain/project"
	"tachyon/internal/client/application/domain/session"
	"tachyon/internal/client/application/domain/user"
	"tachyon/internal/client/remote"
	"tachyon/internal/pkg/event"
	"tachyon/pkg/log"
	"tachyon/pkg/uuid"
)

var ErrSessionNotFound = errors.New("session not found")

type SessionRepository interface {
	Session() (session.Session, error)
	Sessions() ([]session.Session, error)
	ProcessEvents(events ...remote.Event)
}

var ErrProjectNotFound = errors.New("project not found")

type ProjectRepository interface {
	Project() (project.Project, error)
	Projects() ([]project.Project, error)
	ProcessEvents(events ...remote.Event)
}

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	User() (user.User, error)
	Users() ([]user.User, error)
	ProcessEvents(events ...remote.Event)
}

type Application struct {
	remote   *remote.Remote
	sessions SessionRepository
	projects ProjectRepository
	users    UserRepository
}

func New(remote *remote.Remote, sessions SessionRepository, projects ProjectRepository, users UserRepository) (*Application, error) {
	app := &Application{
		remote:   remote,
		sessions: sessions,
		projects: projects,
		users:    users,
	}

	go func() {
		for e := range remote.Events() {
			err := app.handleEvent(e)
			if err != nil {
				log.Errorln(err)
			}
		}
	}()

	return app, nil
}

// func (app *Application) handleConnectedEvent(e event.ConnectedEvent) error {
//	//uid, err := uuid.Parse(e.UserId)
//	//if err != nil {
//	//	return err
//	//}
//
//	//app.user = &user.User{
//	//	Id:   uid,
//	//	Name: uid.String(),
//	//}
//
//	return nil
// }

func (app *Application) CreateSession(name string) error {
	err := app.remote.Publish(event.CreateProjectEvent{
		Name: "my-project",
	})
	if err != nil {
		return err
	}

	time.Sleep(10 * time.Millisecond)

	proj, err := app.projects.Project()
	if err != nil {
		return err
	}

	return app.remote.Publish(event.CreateSessionEvent{
		Name:      name,
		ProjectId: proj.Id.String(),
	})
}

func (app *Application) Project() (project.Project, error) {
	return app.projects.Project()
}

func (app *Application) Session() (session.Session, error) {
	return app.sessions.Session()
}

func (app *Application) handleEvent(e remote.Event) error {
	switch e.AggregateType() {
	case event.User:
		app.users.ProcessEvents(e)
		return nil
	case event.Session:
		app.sessions.ProcessEvents(e)
		return nil
	case event.Project:
		app.projects.ProcessEvents(e)
		return nil
	default:
		return errors.New("invalid aggregate type")
	}
	// switch e := e.(type) {
	// case event.ConnectedEvent:
	//	return app.user.ProcessEvent(e)
	// case event.SessionCreatedEvent:
	//	return app.handleSessionCreatedEvent(e)
	// default:
	//	return errors.New("no event handler")
	// }
}

// func (app *Application) handleSessionCreatedEvent(e event.SessionCreatedEvent) error {
//
// }

func (app *Application) CreatePath(
	projectId uuid.UUID, color project.Color, point project.Vector2,
) error {
	return app.remote.Publish(event.CreatePathEvent{
		ProjectId: projectId.String(),
		Tool:      "pen",
		Color:     color.String(),
		Point: event.Vector2{
			X: point.X,
			Y: point.Y,
		},
	})
}

func (app *Application) CreateProject(name string) error {
	return app.remote.Publish(event.CreateProjectEvent{
		Name: name,
	})
}

// func (app *Application) Project(projectId uuid.UUID) (project.Project, error) {
// 	return app.remote.Project(projectId)
// }
