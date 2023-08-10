package socket

import (
	"context"
	"fmt"
	"time"
)

type Context struct {
	ctx  context.Context
	id   string
	keys map[string]string
	data any
	//router        *Router[T, E]
	subscriptions []func()
}

type HandlerFunc[E any] func(e E)

func (c *Context) String() string {
	return fmt.Sprintf("{ keys: %#v, data: %#v }", c.keys, c.data)
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	if c.ctx == nil {
		return
	}

	return c.ctx.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	if c.ctx == nil {
		return nil
	}

	return c.ctx.Done()
}

func (c *Context) Err() error {
	if c.ctx == nil {
		return nil
	}

	return c.ctx.Err()
}

func (c *Context) Value(key interface{}) interface{} {
	if keyAsString, ok := key.(string); ok {
		if val, exists := c.Get(keyAsString); exists {
			return val
		}
	}

	if c.ctx == nil {
		return nil
	}

	return c.ctx.Value(key)
}

//func (c *Context) Bind(v interface{}) error {
//	return c.router.Unmarshaler(c.data, v)
//}
//
//func (c *Context) Data() []byte {
//	return c.data
//}

func (c *Context) Get(key string) (value string, exists bool) {
	if c.keys == nil {
		return "", false
	}

	value, exists = c.keys[key]

	return
}

func (c *Context) Set(key string, value string) {
	if c.keys == nil {
		c.keys = map[string]string{}
	}

	c.keys[key] = value
}

func (c *Context) Id() string {
	return c.id
}

func (c *Context) Join(channel string) {
	//sub := c.router.pubsub.Subscribe(
	//	channel, c.Id(),
	//)
	//
	//c.subscriptions = append(c.subscriptions, sub)
	//
	//go func() {
	//	for b := range sub.Events() {
	//		//log.Printf("received topic %s, sub %s, payload: %s", channel, sub.Id(), string(b))
	//		if c.Err() != nil {
	//			break
	//		}
	//
	//		err := c.router.EventHandler(c.Id(), b)
	//		if err != nil {
	//			return
	//		}
	//	}
	//}()
}

func (c *Context) send() {
	//b, err := c.router.Marshaler(v)
	//if err != nil {
	//	return
	//}

	//err := c.router.EventHandler(
	//	c.Id(), e,
	//)
	//if err != nil {
	//	return
	//}

	return
}

func (c *Context) Broadcast(topic string, e E) {
	//b, err := c.router.Marshaler(v)
	//if err != nil {
	//	return
	//}

	// dispatch event only on contexes of other clients
	//c.router.pubsub.Publish(
	//	topic, c.Id(), e,
	//)
}

func (c *Context) Send(e E) {
	c.send(e)
}

func (c *Context) Dispose() {
	for _, sub := range c.subscriptions {
		sub.Dispose()
	}
}

//func (c *Context[E]) Handle(e E) {
//	//c.router.dispatcher.Dispatch(c, et, e)
//}
