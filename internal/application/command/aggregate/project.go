package aggregate

import (
	"fmt"

	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/event"
)

type Project struct {
	project.Project
}

func (p *Project) ProcessEvent(e event.Event) {
	switch e := e.(type) {
	case event.ProjectCreatedEvent:
		p.Id = e.Id
		p.Name = e.Name
		p.OwnerId = e.OwnerId
	default:
		fmt.Println("project: unknown event", e)
	}
}
