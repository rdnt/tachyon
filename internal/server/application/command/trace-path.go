package command

import (
	"errors"

	"golang.org/x/exp/slices"

	"tachyon/internal/server/application/domain/project/path"
	"tachyon/internal/server/application/event"
	"tachyon/pkg/uuid"
)

type TracePathArgs struct {
	PathId    uuid.UUID
	UserId    uuid.UUID
	ProjectId uuid.UUID
	Tool      path.Tool
	Color     path.Color
	Point     path.Vector2
}

func (s *service) TracePath(args TracePathArgs) error {
	proj, err := s.projects.Project(args.ProjectId)
	if err != nil {
		return err
	}

	sessions, err := s.sessions.ProjectSessions(proj.Id)
	if err != nil {
		return err
	}

	var found bool
	for _, sess := range sessions {
		idx := slices.Index(sess.UserIds, args.UserId)
		if idx != -1 {
			found = true
		}
	}

	if !found && proj.OwnerId != args.UserId {
		return errors.New("user doesn't have access to the project")
	}

	e := event.PathTracedEvent{
		PathId:    args.PathId,
		UserId:    args.UserId,
		ProjectId: proj.Id,
		Point:     args.Point,
	}

	err = s.publish(e)
	if err != nil {
		return err
	}

	return nil
}
