package command

import (
	"errors"

	"github.com/rdnt/tachyon/internal/server/application/domain/project/path"
	"github.com/rdnt/tachyon/internal/server/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
	"golang.org/x/exp/slices"
)

type CreatePathArgs struct {
	PathId    uuid.UUID
	UserId    uuid.UUID
	ProjectId uuid.UUID
	Tool      path.Tool
	Color     path.Color
	Point     path.Vector2
}

func (s *service) CreatePath(args CreatePathArgs) error {
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

	e := event.PathCreatedEvent{
		PathId:    args.PathId,
		UserId:    args.UserId,
		ProjectId: proj.Id,
		Tool:      args.Tool,
		Color:     args.Color,
		Point:     args.Point,
	}

	err = s.publish(e)
	if err != nil {
		return err
	}

	return nil
}
