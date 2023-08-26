package command

import (
	"fmt"

	"tachyon/internal/server/application/domain/project/path"
	"tachyon/internal/server/application/event"
	"tachyon/pkg/uuid"
)

type CreatePathArgs struct {
	PathId    uuid.UUID
	UserId    uuid.UUID
	ProjectId uuid.UUID
	Tool      path.Tool
	Color     path.Color
	Point     path.Vector2
}

func (s *Commands) CreatePath(args CreatePathArgs) error {
	fmt.Println("CreatePath", args)
	proj, err := s.projects.Project(args.ProjectId)
	if err != nil {
		return err
	}

	// FIXME: limit edit to session users && owner
	//sessions, err := s.sessions.ProjectSessions(proj.Id)
	//if err != nil {
	//	return err
	//}
	//
	//var found bool
	//for _, sess := range sessions {
	//	idx := slices.Index(sess.UserIds, args.UserId)
	//	if idx != -1 {
	//		found = true
	//	}
	//}
	//
	//if !found && proj.OwnerId != args.UserId {
	//	return errors.NewHandler("user doesn't have access to the project")
	//}

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
