package websocket

import (
	"tachyon/internal/pkg/event"
	"tachyon/internal/server/application/command"
	"tachyon/internal/server/application/domain/project/path"
	"tachyon/pkg/uuid"
)

func (s *Server) CreatePath(e event.CreatePathEvent, c *Conn) error {
	uid := c.id

	pid, err := uuid.Parse(e.ProjectId)
	if err != nil {
		return err
	}

	color, err := path.ColorFromString(e.Color)
	if err != nil {
		return err
	}

	coords := path.Vector2{
		X: e.Point.X,
		Y: e.Point.Y,
	}

	pathId := uuid.New()
	tool := path.Pen

	err = s.commands.CreatePath(command.CreatePathArgs{
		PathId:    pathId,
		UserId:    uid,
		ProjectId: pid,
		Tool:      tool,
		Color:     color,
		Point:     coords,
	})
	if err != nil {
		return err
	}

	return s.Publish(pid, event.PathCreatedEvent{
		PathId:    pathId.String(),
		UserId:    uid.String(),
		ProjectId: pid.String(),
		Tool:      tool.String(),
		Color:     color.String(),
		Point: event.Vector2{
			X: e.Point.X,
			Y: e.Point.Y,
		},
	})
}
