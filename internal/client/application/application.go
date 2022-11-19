package application

import (
	"fmt"
	"tachyon/internal/client/application/domain/project"
	"tachyon/internal/client/remote"
	"tachyon/internal/pkg/event"
)

type Application struct {
	remote  *remote.Remote
	project *project.Project
}

func New(remote *remote.Remote) (*Application, error) {
	a := &Application{
		remote: remote,
	}

	go func() {
		for e := range remote.Events() {
			switch e := e.(type) {
			case event.ConnectedEvent:
				fmt.Println(e)
			}
		}
	}()

	return a, nil
}

func (app *Application) CreateSession(name string) error {
	return app.remote.Publish(event.CreateSessionEvent{
		Name: name,
	})
}

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
