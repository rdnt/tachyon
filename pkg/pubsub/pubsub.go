package pubsub

import (
	"fmt"
	"sync"

	"tachyon/pkg/uuid"

	"tachyon/pkg/logger"
)

var log = logger.New("pubsub", logger.MagentaFg)

type MessageBroker[M any] struct {
	mux    sync.Mutex
	wg     sync.WaitGroup
	events map[string]map[uuid.UUID]*Subscription[M]
	closed bool
}

//type PubSub interface {
//	Subscribe(channels ...string) error
//	Unsubscribe(channels ...string) error
//	Publish(channel string, message []byte) error
//}
//
//type Broker interface {
//	Subscribe(sub Subscription, channels ...string) error
//	Unsubscribe(sub Subscription, channels ...string) error
//}

//type Hub interface {
//	// Publish sends input message to specified channels.
//	Publish(channels []string, msg interface{})
//	// Subscribe opens channel to listen specified channels.
//	Subscribe(channels []string) (Channel, error)
//	// Close stops the pubsub hub.
//	Close() error
//}

//type Channel interface {
//	// Events returns channel to receive sessions.
//	Events() <-chan []byte
//	// Close stops listening underlying pubsub topics.
//	Close() error
//	// Done returns channel to receive event when this channel is closed.
//	Done() <-chan bool
//}

// Subscribe creates and returns a new subscription to a topic
func (d *MessageBroker[M]) Subscribe(topic, clientId string) *Subscription[M] {
	d.mux.Lock()
	defer d.mux.Unlock()

	_, ok := d.events[topic]
	if !ok {
		d.events[topic] = make(map[uuid.UUID]*Subscription[M])
	}

	sub := &Subscription[M]{
		dispatcher: d,
		topic:      topic,
		id:         uuid.New(),
		clientId:   clientId,
		events:     make(chan M),
		done:       make(chan bool),
	}

	d.wg.Add(1)
	d.events[topic][sub.id] = sub

	return sub
}

func (d *MessageBroker[M]) Publish(topic, clientId string, e M) {
	d.mux.Lock()
	defer d.mux.Unlock()

	if d.closed {
		return
	}

	subs, ok := d.events[topic]
	if !ok {
		return
	}

	for _, sub := range subs {
		if sub.clientId == clientId {
			continue
		}

		//log.Printf("forwarding to topic %s, sub %s, payload: %s", topic, sub.id, string(b))
		sub.events <- e
	}
}

func (d *MessageBroker[M]) removeSubscription(s *Subscription[M]) {
	subs, ok := d.events[s.topic]
	if !ok {
		return
	}

	sub, ok := subs[s.id]
	if !ok {
		return
	}

	close(sub.done)
	close(sub.events)

	d.wg.Done()
	delete(d.events[s.topic], s.id)
}

func (d *MessageBroker[M]) disposeSubscription(s *Subscription[M]) {
	d.mux.Lock()
	defer d.mux.Unlock()

	d.removeSubscription(s)
}

func (d *MessageBroker[M]) Close() {
	d.mux.Lock()
	defer d.mux.Unlock()

	d.closed = true

	//fmt.Println("disposing subscriptions...")
	for _, subscriptions := range d.events {
		for _, sub := range subscriptions {
			d.removeSubscription(sub)
		}
	}
	//fmt.Println("waiting for consumers...")

	//d.wg.Wait()

	fmt.Println("waiting done")
}

func New[M any]() *MessageBroker[M] {
	return &MessageBroker[M]{
		events: make(map[string]map[uuid.UUID]*Subscription[M]),
	}
}

//// ------------------------------------------------------------

type Subscription[M any] struct {
	mux        sync.Mutex
	dispatcher *MessageBroker[M]
	topic      string
	id         uuid.UUID
	clientId   string
	events     chan M
	done       chan bool
	disposed   bool
}

func (s *Subscription[M]) Dispose() {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.disposed {
		return
	}
	// MUX
	s.disposed = true

	s.dispatcher.disposeSubscription(s)
}

func (s *Subscription[M]) Events() <-chan M {
	return s.events
}

func (s *Subscription[M]) Done() chan bool {
	return s.done
}

func (s *Subscription[M]) Id() uuid.UUID {
	return s.id
}
