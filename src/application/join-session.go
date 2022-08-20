package application

import (
	"fmt"

	"tachyon2/src/application/domain/session"
	"tachyon2/src/application/domain/user"

	"golang.org/x/exp/slices"
)

func (app *App) JoinSession(userId user.Id, sessionId session.Id) (session.Session, error) {
	u, err := app.users.User(userId)
	if err != nil {
		return session.Session{}, err
	}

	s, err := app.sessions.Session(sessionId)
	if err != nil {
		return session.Session{}, err
	}

	s, err = addSessionUser(s, u.Id)
	if err != nil {
		return session.Session{}, err
	}

	s, err = app.sessions.UpdateSession(s)
	if err != nil {
		return session.Session{}, err
	}

	app.events.UserJoinedSession.publish(event.UserJoinedSession{
		SessionId: sessionId,
		UserId:    userId,
	})

	return s, nil
}

func addSessionUser(s session.Session, id user.Id) (session.Session, error) {
	if slices.Contains(s.UserIds, id) {
		return s, fmt.Errorf("user already joined")
	}

	s.UserIds = append(s.UserIds, id)
	return s, nil
}
