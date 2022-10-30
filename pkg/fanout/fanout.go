package fanout

import (
	"fmt"
	"sync"
)

type Exchange[E any] interface {
	Publish(event E) error
	Subscribe() (chan E, func(), error)
}

type FanOut[E any] struct {
	mux           sync.Mutex
	idx           int
	subscriptions map[int]chan E
}

func New[E any]() Exchange[E] {
	return &FanOut[E]{subscriptions: map[int]chan E{}}
}

func (f *FanOut[E]) Subscribe() (chan E, func(), error) {
	f.mux.Lock()
	defer f.mux.Unlock()

	if f.subscriptions == nil {
		f.subscriptions = map[int]chan E{}
	}

	sub := make(chan E)

	idx := f.idx
	f.subscriptions[f.idx] = sub
	f.idx++

	dispose := func() {
		f.mux.Lock()
		delete(f.subscriptions, idx)
		f.mux.Unlock()
	}

	return sub, dispose, nil
}

func (f *FanOut[E]) Publish(event E) error {
	f.mux.Lock()
	subs := f.subscriptions
	f.mux.Unlock()

	for _, conn := range subs {
		conn <- event
	}

	return nil
}

func (f *FanOut[E]) String() string {
	return fmt.Sprintln(len(f.subscriptions), "subscriptions")
}
