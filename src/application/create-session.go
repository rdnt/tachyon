package application

import (
	"tachyon2/src/application/domain/project"
	"tachyon2/src/application/domain/session"
	"tachyon2/src/application/domain/user"
)

func (app *App) CreateSession(
	userId user.Id, projectId project.Id, sessionName string,
) (session.Session, error) {
	u, err := app.users.User(userId)
	if err != nil {
		return session.Session{}, err
	}

	// TODO: remove
	p, err := app.projects.Project("my-project")
	if err != nil {
		p, err = app.projects.CreateProject(
			project.New(sessionName, u.Id),
		)
		if err != nil {
			return session.Session{}, err
		}
	}

	sess, err := app.sessions.CreateSession(
		session.New(sessionName, p.Id, u.Id, u.Id),
	)
	if err != nil {
		return session.Session{}, err
	}

	return sess, nil
}
