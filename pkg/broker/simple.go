package broker

import (
	"sync"

	"github.com/google/uuid"
)

type Simple[E any] struct {
	lock          sync.Mutex
	subscriptions map[string]chan E
}

func NewSimple[E any]() *Simple[E] {
	return &Simple[E]{
		subscriptions: make(map[string]chan E),
	}
}

func (o *Simple[E]) Subscribe() (events chan E, dispose func()) {
	id := uuid.NewString()

	c := make(chan E)

	o.lock.Lock()
	defer o.lock.Unlock()

	if o.subscriptions == nil {
		o.subscriptions = make(map[string]chan E)
	}

	o.subscriptions[id] = c

	return c, func() {
		o.dispose(id)
	}
}

func (o *Simple[E]) Publish(e E) {
	o.lock.Lock()
	defer o.lock.Unlock()

	for _, c := range o.subscriptions {
		go func() {
			c <- e
		}()
	}
}

func (o *Simple[E]) dispose(id string) {
	o.lock.Lock()
	defer o.lock.Unlock()

	if c, ok := o.subscriptions[id]; ok {
		close(c)
	}

	delete(o.subscriptions, id)
}
