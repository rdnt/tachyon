package remote

import (
	"errors"
	"fmt"
	"log"

	"github.com/gorilla/websocket"

	"tachyon/internal/pkg/event"
)

type Remote struct {
	address string
	conn    *websocket.Conn
	events  chan event.Event
}

func New(address string) (*Remote, error) {
	r := &Remote{
		address: address,
		events:  make(chan event.Event),
	}

	conn, _, err := websocket.DefaultDialer.Dial(address, nil)
	if err != nil {
		return nil, err
	}

	r.conn = conn

	go func() {
		for {
			typ, b, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}

			if typ != websocket.TextMessage {
				fmt.Println("invalid payload type")
				continue
			}

			e, err := event.FromJSON(b)
			if err != nil {
				fmt.Println("invalid message type", err)
				continue
			}

			r.events <- e
		}
	}()

	return r, nil
}

func (r *Remote) Publish(e event.Event) error {
	if r.conn == nil {
		return errors.New("connection not established")
	}

	b, err := event.ToJSON(e)
	if err != nil {
		return err
	}

	err = r.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		return err
	}

	return nil
}

func (r *Remote) Events() chan event.Event {
	return r.events
}

func (r *Remote) HandleConnectedEvent(e event.ConnectedEvent) {
	fmt.Println("CONNECTED", e)
}
