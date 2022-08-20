package user

import (
	"fmt"

	"tachyon2/pkg/uuid"
)

type Id string

type User struct {
	Id      Id
	Name    string
	Pointer Pointer
}

func (u User) Subscribe() {

}

func (u *User) String() string {
	return fmt.Sprintf("User{id: %s, name: %s, pointer: %v}", u.Id, u.Name, u.Pointer)
}

func New() User {
	id := uuid.New()

	u := User{
		Id: Id(id),
		Pointer: Pointer{
			Mode: Hover,
			X:    0,
			Y:    0,
		},
	}

	return u
}
