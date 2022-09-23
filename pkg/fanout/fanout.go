package fanout

import (
	"fmt"
	"sync"
)

type Fanout[E any] struct {
	lock  sync.Mutex
	conns []chan E
}

func New[E any]() *Fanout[E] {
	return &Fanout[E]{conns: []chan E{}}
}

func (f *Fanout[E]) Subscribe() (chan E, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.conns == nil {
		f.conns = []chan E{}
	}

	conn := make(chan E)

	f.conns = append(f.conns, conn)

	return conn, nil
}

func (f *Fanout[E]) Publish(event E) error {
	f.lock.Lock()
	conns := f.conns
	f.lock.Unlock()

	for _, conn := range conns {
		conn <- event
	}

	return nil
}

func (f *Fanout[E]) String() string {
	return fmt.Sprintln(len(f.conns), "connections")
}

type exchange[E any] interface {
	Publish(event E) error
	Subscribe() (chan E, error)
}
