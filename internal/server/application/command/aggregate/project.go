package aggregate

import (
	"fmt"

	"tachyon/internal/server/application/domain/project"
	"tachyon/internal/server/application/event"
)

type Project struct {
	project.Project
}

func (p *Project) ProcessEvent(e event.Event) {
	switch e := e.(type) {
	case event.ProjectCreatedEvent:
		p.Id = e.ProjectId
		p.Name = e.Name
		p.OwnerId = e.OwnerId
	case event.PathCreatedEvent:
		//idx := slices.IndexFunc(p.Paths, func(p path.Path) bool {
		//	return p.Coords.X == e.Coords.X && p.Coords.Y == e.Coords.Y
		//})
		//
		//if idx == -1 {
		//	p.Pixels = append(p.Pixels, project.Pixel{
		//		Color:  e.Color,
		//		Coords: e.Coords,
		//	})
		//} else {
		//	p.Pixels[idx] = project.Pixel{
		//		Color:  e.Color,
		//		Coords: e.Coords,
		//	}
		//}
	case event.PathTracedEvent:
		//idx := slices.IndexFunc(p.Pixels, func(pix project.Pixel) bool {
		//	return pix.Coords.X == e.Coords.X && pix.Coords.Y == e.Coords.Y
		//})
		//
		//if idx != -1 {
		//	p.Pixels = slices.Delete(p.Pixels, idx, idx+1)
		//}
	default:
		fmt.Println("project: unknown event", e)
	}
}
