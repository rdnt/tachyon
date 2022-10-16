package command

import (
	"fmt"

	"github.com/rdnt/tachyon/internal/application/domain/project/path"
	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
	"golang.org/x/exp/slices"
)

type CreatePathArgs struct {
	PathId    path.Id
	UserId    user.Id
	SessionId session.Id
	Tool      path.Tool
	Color     path.Color
	Point     path.Vector2
}

func (s *service) CreatePath(args CreatePathArgs) error {
	sess, err := s.sessions.Session(args.SessionId)
	if err != nil {
		return err
	}

	idx := slices.Index(sess.UserIds, args.UserId)
	if idx == -1 {
		return fmt.Errorf("cannot remove user from session: not a member")
	}

	e := event.NewPathCreatedEvent(event.PathCreatedEvent{
		PathId:    args.PathId,
		UserId:    args.UserId,
		ProjectId: sess.ProjectId,
		Tool:      args.Tool,
		Color:     args.Color,
		Point:     args.Point,
	})

	err = s.publish(e)
	if err != nil {
		return err
	}

	return nil
}
