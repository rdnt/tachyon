package user

import (
	"fmt"

	"github.com/rdnt/tachyon/pkg/uuid"
)

type Id uuid.UUID

type User struct {
	Id   Id
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("User{id: %s, name: %s}", u.Id, u.Name)
}
