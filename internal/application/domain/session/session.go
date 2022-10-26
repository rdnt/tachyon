package session

import (
	"github.com/rdnt/tachyon/pkg/uuid"
)

type Session struct {
	Id        uuid.UUID
	Name      string
	ProjectId uuid.UUID
	UserIds   []uuid.UUID
}
