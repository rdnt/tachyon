package session

import (
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type Id uuid.UUID

type Session struct {
	Id        Id
	Name      string
	ProjectId project.Id
	UserIds   []user.Id
}
