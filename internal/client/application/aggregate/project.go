package aggregate

import (
	"fmt"

	"tachyon/internal/server/application/domain/project/path"

	"tachyon/internal/client/application/domain/project"
	"tachyon/internal/pkg/event"
	"tachyon/pkg/uuid"
)

type Project struct {
	project.Project
}

func (p *Project) ProcessEvent(e event.Event) {
	switch e := e.(type) {
	case event.ProjectCreatedEvent:
		p.Project.Id = uuid.MustParse(e.ProjectId)
	case event.PathCreatedEvent:
		//idx := slices.IndexFunc(p.Pixels, func(pix project.Pixel) bool {
		//	return pix.Coords.X == e.Coords.X && pix.Coords.Y == e.Coords.Y
		//})
		clr, err := path.ColorFromString("#ffffff")
		if err != nil {
			panic(err)
		}

		p.Project.Paths = append(p.Project.Paths, project.Path{
			Id:    uuid.MustParse(e.PathId),
			Tool:  e.Tool,
			Color: project.Color(clr),
			//Color:  parseColor(e.Color),
			Points: []project.Vector2{
				{
					X: e.Point.X,
					Y: e.Point.Y,
				},
			},
		})
	//case event.PixelErasedEvent:
	//	idx := slices.IndexFunc(p.Pixels, func(pix project.Pixel) bool {
	//		return pix.Coords.X == e.Coords.X && pix.Coords.Y == e.Coords.Y
	//	})
	//
	//	if idx != -1 {
	//		p.Pixels = slices.Delete(p.Pixels, idx, idx+1)
	//	}
	default:
		fmt.Println("project: unknown event", e)
	}
}
