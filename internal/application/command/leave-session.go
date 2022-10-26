package command

import (
	"fmt"

	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
	"golang.org/x/exp/slices"
)

func (s *service) LeaveSession(id uuid.UUID, uid uuid.UUID) error {
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
