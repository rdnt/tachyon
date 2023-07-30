package broker

import (
	"sync"

	"tachyon/pkg/uuid"
)

type Broker[C comparable, E any] struct {
	lock          sync.Mutex
	subscriptions map[C]map[uuid.UUID]func(E)
}

func New[C comparable, E any]() *Broker[C, E] {
	return &Broker[C, E]{
		subscriptions: map[C]map[uuid.UUID]func(E){},
	}
}

func (o *Broker[C, E]) Subscribe(channel C, handler func(e E)) (dispose func()) {
	id := uuid.New()

	o.lock.Lock()
	defer o.lock.Unlock()

	if _, ok := o.subscriptions[channel]; !ok {
		o.subscriptions[channel] = map[uuid.UUID]func(E){}
	}

	o.subscriptions[channel][id] = handler

	return func() {
		o.dispose(channel, id)
	}
}

func (o *Broker[C, E]) Publish(channel C, e E) {
	o.lock.Lock()
	defer o.lock.Unlock()

	for ch, subs := range o.subscriptions {
		if ch != channel {
			continue
		}

		for _, h := range subs {
			if h != nil {
				go func(h func(E)) {
					// TODO: remove simulated network delay
					// time.Sleep(10 * time.Millisecond)
					h(e)
				}(h)
			}
		}
	}
}

func (o *Broker[C, E]) dispose(channel C, id uuid.UUID) {
	o.lock.Lock()
	defer o.lock.Unlock()

	if _, ok := o.subscriptions[channel]; !ok {
		return
	}

	delete(o.subscriptions[channel], id)
}
