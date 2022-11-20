package application

import (
	"errors"

	"tachyon/internal/client/application/aggregate"
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
	Session(id uuid.UUID) (session.Session, error)
	ProcessEvents(events ...remote.Event)
}

var ErrProjectNotFound = errors.New("project not found")

type ProjectRepository interface {
	Project(id uuid.UUID) (project.Project, error)
	ProcessEvents(events ...remote.Event)
}

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	User(id uuid.UUID) (user.User, error)
	ProcessEvents(events ...remote.Event)
}

type Application struct {
	remote   *remote.Remote
	sessions SessionRepository
	projects ProjectRepository
	users    UserRepository

	session *aggregate.Session
	project *aggregate.Project
	user    *aggregate.User
}

func New(remote *remote.Remote, sessions SessionRepository, projects ProjectRepository, users UserRepository) (*Application, error) {
	app := &Application{
		remote:   remote,
		sessions: sessions,
		projects: projects,
		users:    users,

		project: &aggregate.Project{},
		session: &aggregate.Session{},
		user:    &aggregate.User{},
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

//func (app *Application) handleConnectedEvent(e event.ConnectedEvent) error {
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
//}

func (app *Application) CreateSession(name string) error {
	return app.remote.Publish(event.CreateSessionEvent{
		Name: name,
	})
}

func (app *Application) handleEvent(e remote.Event) error {
	switch e.AggregateType() {
	case event.User:
		app.users.ProcessEvents(e)
		app.user.ProcessEvent(e)
		return nil
	case event.Session:
		app.sessions.ProcessEvents(e)
		app.session.ProcessEvent(e)
		return nil
	case event.Project:
		app.projects.ProcessEvents(e)
		app.project.ProcessEvent(e)
		return nil
	default:
		return errors.New("invalid aggregate type")
	}
	//switch e := e.(type) {
	//case event.ConnectedEvent:
	//	return app.user.ProcessEvent(e)
	//case event.SessionCreatedEvent:
	//	return app.handleSessionCreatedEvent(e)
	//default:
	//	return errors.New("no event handler")
	//}
}

//func (app *Application) handleSessionCreatedEvent(e event.SessionCreatedEvent) error {
//
//}

// func (app *Application) DrawPixel(
// 	projectId uuid.UUID, color project.Color, coords project.Vector2,
// ) error {
// 	return app.remote.Publish(event.UpdatePixelEvent{
// 		ProjectId: projectId.String(),
// 		Color:     color,
// 		Coords:    coords,
// 	})
// }
//
// func (app *Application) Project(projectId uuid.UUID) (project.Project, error) {
// 	return app.remote.Project(projectId)
// }
