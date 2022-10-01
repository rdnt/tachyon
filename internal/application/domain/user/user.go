package user

import (
	"fmt"

	"github.com/google/uuid"
)

type Id uuid.UUID

func (id Id) String() string {
	return uuid.UUID(id).String()
}

type User struct {
	Id   Id
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("User{id: %s, name: %s}", u.Id, u.Name)
}
