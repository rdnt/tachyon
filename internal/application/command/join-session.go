package command

import (
	"fmt"

	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
	"golang.org/x/exp/slices"
)

func (s *service) JoinSession(id session.Id, uid user.Id) error {
	_, err := s.users.User(uid)
	if err != nil {
		return err
	}

	sess, err := s.sessions.Session(id)
	if err != nil {
		return err
	}

	if slices.Index(sess.UserIds, uid) != -1 {
		return fmt.Errorf("cannot add user to session: already a member")
	}

	e := event.NewJoinedSessionEvent(event.JoinedSessionEvent{
		SessionId: sess.Id,
		UserId:    uid,
	})

	err = s.publish(e)
	if err != nil {
		return err
	}

	return nil
}
