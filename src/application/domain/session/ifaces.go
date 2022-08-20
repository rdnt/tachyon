package session

type Repository interface {
	CreateSession(Session) (Session, error)
	Session(Id) (Session, error)
	UpdateSession(Session) (Session, error)
	Sessions() (map[Id]Session, error)
	DeleteSession(Id) error
}
