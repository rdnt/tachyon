package command

import (
	"fmt"

	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
	"golang.org/x/exp/slices"
)

func (s *service) LeaveSession(id session.Id, uid user.Id) error {
	_, err := s.users.User(uid)
	if err != nil {
		return err
	}

	sess, err := s.sessions.Session(id)
	if err != nil {
		return err
	}

	idx := slices.Index(sess.UserIds, uid)
	if idx == -1 {
		return fmt.Errorf("cannot remove user from session: not a member")
	}

	e := event.NewLeftSessionEvent(event.LeftSessionEvent{
		SessionId: sess.Id,
		UserId:    uid,
	})

	err = s.publish(e)
	if err != nil {
		return err
	}

	return nil
}
