package command

import (
	"fmt"

	"golang.org/x/exp/slices"

	"tachyon/internal/server/application/event"
	"tachyon/pkg/uuid"
)

func (s *Commands) LeaveSession(id uuid.UUID, uid uuid.UUID) error {
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

	e := event.LeftSessionEvent{
		SessionId: sess.Id,
		UserId:    uid,
	}

	err = s.publish(e)
	if err != nil {
		return err
	}

	return nil
}
