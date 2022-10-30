package interfaces

//type EventStore[E any] interface {
//	Publish(E) error
//	Subscribe() (events chan E, dispose func(), err error)
//	Events() ([]E, error)
//}

type EventStore[E any] interface {
	Publish(event E) (err error)
	Subscribe(handler func(E)) (dispose func(), err error)
	Events() (events []E, err error)
}

type EventBus[E any] interface {
	Publish(event E) (err error)
	Subscribe(handler func(E)) (dispose func(), err error)
	Events() (events []E, err error)
}
