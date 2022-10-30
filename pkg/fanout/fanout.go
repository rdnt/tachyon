package fanout

import (
	"fmt"
	"sync"
)

type Exchange[E any] interface {
	Publish(event E) error
	Subscribe(handler func(E)) (func(), error)
	Events() ([]E, error)
}

type Fanout[E any] struct {
	mux       sync.Mutex
	events    []E
	handlerId int
	handlers  map[int]func(E)
}

func New[E any]() Exchange[E] {
	return &Fanout[E]{
		events:   make([]E, 0),
		handlers: make(map[int]func(E), 0),
	}
}

func (f *Fanout[E]) Events() ([]E, error) {
	f.mux.Lock()
	defer f.mux.Unlock()

	return f.events, nil
}

func (f *Fanout[E]) Subscribe(handle func(E)) (func(), error) {
	f.mux.Lock()
	defer f.mux.Unlock()

	id := f.handlerId
	f.handlerId++

	f.handlers[id] = handle

	dispose := func() {
		f.mux.Lock()
		delete(f.handlers, id)
		f.mux.Unlock()
	}

	return dispose, nil
}

func (f *Fanout[E]) Publish(e E) error {
	f.mux.Lock()
	defer f.mux.Unlock()

	for _, handle := range f.handlers {
		handle(e)
	}

	return nil
}

func (f *Fanout[E]) String() string {
	return fmt.Sprintln(len(f.handlers), "handlers", len(f.events), "events")
}
