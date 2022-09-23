package user

import "fmt"

type Id string

type User struct {
	Id   Id
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("User{id: %s, name: %s}", u.Id, u.Name)
}
