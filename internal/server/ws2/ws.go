package ws2

import (
	"fmt"

	"tachyon/pkg/socket"
)

type EventType string
type Event interface {
	EventType() EventType
}

var Open EventType = "open"
var Close EventType = "close"

func New() {
	s := socket.New[EventType, Event]()

	s.Marshaler = func(e Event) ([]byte, error) {
		return []byte(fmt.Sprint(e)), nil
	}

	s.Unmarshaler = func(b []byte, v Event) error {
		//*v = Event(b)
		return nil
	}

	s.EventHandler = func(clientId string, e Event) error {
		fmt.Println("Event handler called", clientId, e)
		return nil
	}

	s.On(Open, func(c *socket.Context, e Event) {
		fmt.Println("new client!")

	})

	s.OnOpen(func(c *socket.Context) {
		c.Broadcast("some-topic", int64(123))
		fmt.Println("client left!")
	})

	s.OnClose(, func(e Event) {
		fmt.Println("SOME EVENT!", e)
	})
}
