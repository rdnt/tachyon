package aggregate

import (
	"fmt"

	"github.com/rdnt/tachyon/internal/server/application/domain/user"
	"github.com/rdnt/tachyon/internal/server/application/event"
)

type User struct {
	user.User
}

func (u *User) ProcessEvent(e event.Event) {
	switch e := e.(type) {
	case event.UserCreatedEvent:
		u.Id = e.UserId
		u.Name = e.Name
	default:
		fmt.Println("user: unknown event", e)
	}
}
