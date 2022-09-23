package session

import (
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/user"
)

type Id string

type Session struct {
	Id        Id
	Name      string
	ProjectId project.Id
	OwnerId   user.Id
	UserIds   []user.Id
}
