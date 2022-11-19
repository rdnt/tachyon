package user_view

import (
	"encoding/json"
	"errors"

	"tachyon/internal/server/application/domain/user"
	"tachyon/pkg/uuid"
)

type View struct {
	users map[uuid.UUID]user.User
}

func (v *View) User(id uuid.UUID) (user.User, error) {
	u, ok := v.users[id]
	if !ok {
		return user.User{}, errors.New("not found")
	}

	return u, nil
}

func (v *View) CreateUser(u user.User) error {
	v.users[u.Id] = u
	return nil
}

func (v *View) String() string {
	b, err := json.Marshal(v.users)
	if err != nil {
		return "error"
	}

	return string(b)
}

func New() *View {
	r := &View{
		users: map[uuid.UUID]user.User{},
	}

	return r
}
