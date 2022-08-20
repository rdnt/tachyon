package application

import (
	"fmt"

	"tachyon2/src/application/domain/session"
	"tachyon2/src/application/domain/user"

	"golang.org/x/exp/slices"
)

func (app *App) LeaveSession(userId user.Id, sessionId session.Id) error {
	u, err := app.users.User(userId)
	if err != nil {
		return err
	}

	s, err := app.sessions.Session(sessionId)
	if err != nil {
		return err
	}

	s, err = removeSessionUser(s, u.Id)
	if err != nil {
		return err
	}

	s, err = app.sessions.UpdateSession(s)
	if err != nil {
		return err
	}

	app.events.UserLeftSession.publish(event.UserLeftSession{
		SessionId: sessionId,
		UserId:    userId,
	})

	return nil
}

func removeSessionUser(s session.Session, id user.Id) (session.Session, error) {
	idx := slices.Index(s.UserIds, id)
	if idx > -1 {
		if id == s.OwnerId {
			s.OwnerId = s.UserIds[(idx+1)%len(s.UserIds)]
		}

		slices.Delete(s.UserIds, idx, idx)

		return s, nil
	}

	return s, fmt.Errorf("user not in session")
}
