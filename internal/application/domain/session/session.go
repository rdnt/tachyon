package session

import (
	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/user"
)

type Id uuid.UUID

type Session struct {
	Id        Id
	Name      string
	ProjectId project.Id
	UserIds   []user.Id
}
