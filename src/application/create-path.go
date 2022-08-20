package application

import (
	"tachyon2/src/application/domain/project"
	"tachyon2/src/application/domain/project/path"
	"tachyon2/src/application/domain/session"
	"tachyon2/src/application/domain/user"
	"tachyon2/src/application/event"
)

func (app *App) CreatePath(userId user.Id, sessionId session.Id, tool path.Tool, color path.Color, x, y float64) (path.Path, error) {
	u, err := app.users.User(userId)
	if err != nil {
		return path.Path{}, err
	}

	sess, err := app.sessions.Session(sessionId)
	if err != nil {
		return path.Path{}, err
	}

	proj, err := app.projects.Project(sess.ProjectId)
	if err != nil {
		return path.Path{}, err
	}

	pth := path.New(tool, color, x, y)

	proj = addProjectUserPath(proj, u.Id, pth)

	_, err = app.projects.UpdateProject(proj)
	if err != nil {
		return path.Path{}, err
	}

	app.events.PathCreated.publish(event.PathCreatedEvent{
		Event: event.Event{AggregateId: app.id}
		SessionId: sessionId,
		UserId:    userId,
		Path:      pth,
	})

	return pth, nil
}

func addProjectUserPath(proj project.Project, userId user.Id, p path.Path) project.Project {
	proj.Paths = append(proj.Paths, p)
	proj.UserHistory[userId] = append(proj.UserHistory[userId], p.Id)
	proj.UserHistoryIndicator[userId]++

	// TODO: actually prune history here

	return proj
}
