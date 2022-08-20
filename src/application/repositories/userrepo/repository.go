package userrepo

import (
	"errors"
	"fmt"
	"sync"

	"tachyon2/pkg/logger"
	"tachyon2/src/application/domain/user"
)

var ErrNotFound = errors.New("user not found")
var ErrExists = errors.New("user already exists")

type Repository struct {
	log   *logger.Logger
	mux   sync.Mutex
	users map[user.Id]user.User
}

func (r *Repository) CreateUser(u user.User) (user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	r.users[u.Id] = u

	r.log.Println("user created:", u)
	r.log.Println(r)

	return u, nil
}

func (r *Repository) User(id user.Id) (user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	u, ok := r.users[id]
	if !ok {
		return user.User{}, ErrNotFound
	}

	return u, nil
}

func (r *Repository) Users() (map[user.Id]user.User, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.users, nil
}

func (r *Repository) DeleteUser(id user.Id) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	delete(r.users, id)

	r.log.Println("user deleted:", id)
	r.log.Println(r)

	return nil
}

func (r *Repository) String() string {
	return fmt.Sprintf("=== %v", r.users)
}

func New() *Repository {
	return &Repository{
		users: map[user.Id]user.User{},
		log:   logger.New("users", logger.RedFg),
	}
}
