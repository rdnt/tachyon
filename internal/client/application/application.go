package application

import (
	"github.com/rdnt/tachyon/internal/client/application/domain/project"
	"github.com/rdnt/tachyon/internal/client/application/event"
	"github.com/rdnt/tachyon/internal/client/remote"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type Application struct {
	remote *remote.Remote
}

func New(remote *remote.Remote) (*Application, error) {
	return &Application{
		remote: remote,
	}, nil
}

func (app *Application) CreateUser(name string) error {
	return app.remote.Publish(event.CreateUserEvent{
		Name: name,
	})
}

func (app *Application) CreateProject(name string) error {
	return app.remote.Publish(event.CreateProjectEvent{
		Name: name,
	})
}

func (app *Application) DrawPixel(
	projectId uuid.UUID, color project.Color, coords project.Vector2,
) error {
	return app.remote.Publish(event.DrawPixelEvent{
		ProjectId: projectId,
		Color:     color,
		Coords:    coords,
	})
}

func (app *Application) Project(projectId uuid.UUID) (project.Project, error) {
	return app.remote.Project(projectId)
}
