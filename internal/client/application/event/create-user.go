package event

const (
	CreateUser Type = "create_user"
)

type CreateUserEvent struct {
	Name string
}

func (CreateUserEvent) Type() Type {
	return CreateUser
}
