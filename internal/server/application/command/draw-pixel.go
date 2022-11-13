package command

import (
	"errors"

	"github.com/rdnt/tachyon/internal/server/application/domain/project"
	"github.com/rdnt/tachyon/internal/server/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
	"golang.org/x/exp/slices"
)

type DrawPixelArgs struct {
	UserId    uuid.UUID
	ProjectId uuid.UUID
	Color     project.Color
	Coords    project.Vector2
}

func (s *service) DrawPixel(args DrawPixelArgs) error {
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

	e := event.PixelDrawnEvent{
		UserId:    args.UserId,
		ProjectId: proj.Id,
		Color:     args.Color,
		Coords:    args.Coords,
	}

	err = s.publish(e)
	if err != nil {
		return err
	}

	return nil
}
