package user

import (
	"fmt"

	"tachyon/pkg/uuid"
)

type User struct {
	Id   uuid.UUID
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("User{id: %s, name: %s}", u.Id, u.Name)
}
