package application

import (
	"fmt"

	"tachyon2/src/application/domain/project/path"
	"tachyon2/src/application/domain/session"
	"tachyon2/src/application/domain/user"
)

func (app *App) TracePath(userId user.Id, sessionId session.Id, pathId path.Id, x, y float64) (path.Path, error) {
	_, err := app.users.User(userId)
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

	var found bool
	var pth path.Path
	for i, p := range proj.Paths {
		if p.Id == pathId {
			p = addPathPoint(p, x, y)
			proj.Paths[i] = p
			pth = p

			found = true
			break
		}
	}

	if !found {
		return path.Path{}, fmt.Errorf("path not found")
	}

	_, err = app.projects.UpdateProject(proj)
	if err != nil {
		return path.Path{}, err
	}

	return pth, nil
}

func addPathPoint(p path.Path, x, y float64) path.Path {
	p.Points = append(p.Points, path.Point{
		X: x,
		Y: y,
	})

	return p
}
