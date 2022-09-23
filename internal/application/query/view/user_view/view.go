package user_view

import (
	"encoding/json"
	"errors"

	"github.com/rdnt/tachyon/internal/application/domain/user"
)

type View struct {
	users map[user.Id]user.User
}

func (v *View) User(id user.Id) (user.User, error) {
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
		users: map[user.Id]user.User{},
	}

	return r
}
