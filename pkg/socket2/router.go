package socket2

import "context"

type Router struct {
	dispatcher *Dispatcher
	pubsub     *pubsub.MessageBroker
	encode     Encoder
	decode     Decoder
	handle     EventHandler
}

func New() *Router {
	r := &Router{
		dispatcher: NewDispatcher(),
		pubsub:     pubsub.New(),
	}

	return r
}

func (r *Router) On(event string, h HandlerFunc) {
	r.dispatcher.On(event, h)
}

func (r *Router) SetEncoder(enc Encoder) {
	r.encode = enc
}

func (r *Router) SetDecoder(dec Decoder) {
	r.decode = dec
}

func (r *Router) SetEventHandler(h EventHandler) {
	r.handle = h
}

func (r *Router) NewContext(ctx context.Context, clientId string) *Context {
	return &Context{
		router:        r,
		id:            clientId,
		ctx:           ctx,
		subscriptions: []*pubsub.Subscription{},
	}
}
