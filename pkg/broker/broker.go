package broker

import (
	"sync"

	"github.com/google/uuid"
)

type Broker[E any] struct {
	lock          sync.Mutex
	subscriptions map[string]func(E)
}

func New[E any]() *Broker[E] {
	return &Broker[E]{
		subscriptions: map[string]func(E){},
	}
}

func (o *Broker[E]) Subscribe(handler func(e E)) (dispose func()) {
	o.lock.Lock()
	defer o.lock.Unlock()

	id := uuid.New().String()
	o.subscriptions[id] = handler

	return func() {
		o.dispose(id)
	}
}

func (o *Broker[E]) Publish(e E) {
	o.lock.Lock()
	defer o.lock.Unlock()

	for _, h := range o.subscriptions {
		if h != nil {
			h(e)
		}
	}
}

func (o *Broker[E]) dispose(id string) {
	o.lock.Lock()
	defer o.lock.Unlock()

	delete(o.subscriptions, id)
}
