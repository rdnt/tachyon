package command

import (
	"errors"

	"github.com/rdnt/tachyon/internal/server/application/domain/project"
	"github.com/rdnt/tachyon/internal/server/application/event"
	"github.com/rdnt/tachyon/pkg/uuid"
	"golang.org/x/exp/slices"
)

type ErasePixelArgs struct {
	UserId    uuid.UUID
	ProjectId uuid.UUID
	Coords    project.Vector2
}

func (s *service) ErasePixel(args ErasePixelArgs) error {
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

	//for _, pix := range proj.Pixels {
	//	if pix.Coords.X == args.Coords.X && pix.Coords.Y == args.Coords.Y {
	//		idx := slices.IndexFunc(proj.Pixels, func(p project.Pixel) bool {
	//			return pix.Coords.X == args.Coords.X && pix.Coords.Y == args.Coords.Y
	//		})
	//
	//		proj.Pixels = slices.Delete(proj.Pixels, idx, idx+1)
	//	}
	//}

	e := event.PixelErasedEvent{
		UserId:    args.UserId,
		ProjectId: proj.Id,
		Coords:    args.Coords,
	}

	err = s.publish(e)
	if err != nil {
		return err
	}

	return nil
}
