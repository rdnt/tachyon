package fanout

import (
	"fmt"
	"sync"
)

type Exchange[E any] interface {
	Publish(event E) error
	Subscribe() (chan E, error)
}

type FanOut[E any] struct {
	lock          sync.Mutex
	subscriptions []chan E
}

func New[E any]() Exchange[E] {
	return &FanOut[E]{subscriptions: []chan E{}}
}

func (f *FanOut[E]) Subscribe() (chan E, error) {
	f.lock.Lock()
	defer f.lock.Unlock()

	if f.subscriptions == nil {
		f.subscriptions = []chan E{}
	}

	sub := make(chan E)

	f.subscriptions = append(f.subscriptions, sub)

	return sub, nil
}

func (f *FanOut[E]) Publish(event E) error {
	f.lock.Lock()
	subs := f.subscriptions
	f.lock.Unlock()

	for _, conn := range subs {
		conn <- event
	}

	return nil
}

func (f *FanOut[E]) String() string {
	return fmt.Sprintln(len(f.subscriptions), "subscriptions")
}
