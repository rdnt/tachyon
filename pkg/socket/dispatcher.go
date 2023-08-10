package socket

import (
	"fmt"
)

type Dispatcher[T comparable, E any] struct {
	events map[T][]HandlerFunc[E]
}

func NewDispatcher[T comparable, E any]() *Dispatcher[T, E] {
	return &Dispatcher[T, E]{
		events: make(map[T][]HandlerFunc[E], 0),
	}
}

func (d *Dispatcher[T, E]) On(name T, h HandlerFunc[E]) {
	_, ok := d.events[name]
	if !ok {
		d.events[name] = make([]HandlerFunc[E], 0, 1)
	}

	d.events[name] = append(d.events[name], h)

	//log.Printf("added listener for event %s", name)
}

func (d *Dispatcher[T, E]) Dispatch(c *Context[T, E], name T, e E) {
	handlers, ok := d.events[name]
	if !ok {
		fmt.Printf("%s event is not registered\n", name)
		return
	}

	c.data = e

	//log.Printf("triggering event %s, data %s", name, string(b))

	for _, h := range handlers {
		h(e)
	}
}
