package interfaces

type EventStore[E any] interface {
	Publish(E) error
	Subscribe() (events chan E, dispose func(), err error)
	Events() ([]E, error)
}
