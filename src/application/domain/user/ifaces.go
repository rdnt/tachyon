package user

type Repository interface {
	CreateUser(User) (User, error)
	User(Id) (User, error)
	Users() (map[Id]User, error)
	DeleteUser(Id) error
}
