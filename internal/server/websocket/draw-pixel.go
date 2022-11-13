package websocket

import (
	wsevent "github.com/rdnt/tachyon/cmd/server/websocket/event"
	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/pkg/uuid"
)

func (s *Server) DrawPixel(e wsevent.DrawPixelEvent, c *Conn) error {
	uid, err := uuid.Parse(c.Get("userId"))
	if err != nil {
		return err
	}

	pid, err := uuid.Parse(e.ProjectId)
	if err != nil {
		return err
	}

	color, err := project.ColorFromString(e.Color)
	if err != nil {
		return err
	}

	coords := project.Vector2{
		X: e.Coords.X,
		Y: e.Coords.Y,
	}

	err = s.commands.DrawPixel(command.DrawPixelArgs{
		UserId:    uid,
		ProjectId: pid,
		Color:     color,
		Coords:    coords,
	})
	if err != nil {
		return err
	}

	//b, err := wsevent.PixelDrawnEventToJSON(wsevent.PixelDrawnEvent{
	//	UserId:    uid.String(),
	//	ProjectId: pid.String(),
	//	Color:     color.String(),
	//	Coords:    e.Coords,
	//})
	//if err != nil {
	//	return err
	//}
	//
	//return c.Write(b)
	return nil
}
