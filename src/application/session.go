package application

import "tachyon2/src/application/domain/session"

func (app *App) Session(id session.Id) (session.Session, error) {
	return app.sessions.Session(id)
}
