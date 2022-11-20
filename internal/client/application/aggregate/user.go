package aggregate

import (
	"fmt"

	"tachyon/internal/pkg/event"
	"tachyon/pkg/uuid"

	"tachyon/internal/client/application/domain/user"
)

type User struct {
	user.User
}

func (u *User) ProcessEvent(e event.Event) {
	switch e := e.(type) {
	case event.ConnectedEvent:
		u.Id = uuid.MustParse(e.UserId)
		u.Name = e.UserId
	default:
		fmt.Println("user: unknown event", e)
	}
}
