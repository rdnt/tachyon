package socket

import (
	"context"

	"tachyon/pkg/pubsub"
)

var (
	Connected    = "connected"
	Disconnected = "disconnected"
)

type Client interface {
	Id() string
	Send(b []byte) error
	Receive() ([]byte, error)
}

type Marshaler[E any] func(E) ([]byte, error)

type Unmarshaler[E any] func(b []byte, v E) error

type EventHandler[T comparable, E any] func(clientId string, e E) error

type Router[T comparable, E any] struct {
	//dispatcher   *Dispatcher[T, E]
	pubsub       *pubsub.MessageBroker[E]
	Marshaler    Marshaler[E]
	Unmarshaler  Unmarshaler[E]
	EventHandler EventHandler[T, E]
}

func New[T comparable, E any]() *Router[T, E] {
	r := &Router[T, E]{
		//dispatcher: NewDispatcher[T, E](),
		pubsub: pubsub.New[E](),
	}

	return r
}

func (r *Router[T, E]) On(event T, h func(c *Context[T, E], e E)) {
	//r.dispatcher.On(event, h)
}

func (r *Router[T, E]) OnOpen(h func(c *Context[T, E])) {
	//r.dispatcher.On(event, h)
}

func (r *Router[T, E]) OnClose(h func(c *Context[T, E])) {
	//r.dispatcher.On(event, h)
}

func (r *Router[T, E]) NewContext(ctx context.Context, clientId string) *Context[T, E] {
	return &Context[T, E]{
		//router:        r,
		id:            clientId,
		ctx:           ctx,
		subscriptions: []*pubsub.Subscription[E]{},
	}
}
