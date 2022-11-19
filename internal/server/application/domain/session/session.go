package session

import (
	"tachyon/pkg/uuid"
)

type Session struct {
	Id        uuid.UUID
	Name      string
	ProjectId uuid.UUID
	UserIds   []uuid.UUID
}
