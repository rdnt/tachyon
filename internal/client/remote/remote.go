package remote

import (
	"errors"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/rdnt/tachyon/internal/client/application/event"
	"github.com/rdnt/tachyon/internal/client/remote/websocketevent"
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
				fmt.Println("invalid message type")
				continue
			}

			e, err := websocketevent.FromJSON(b)
			if err != nil {
				fmt.Println(err)
				continue
			}

			r.handleEvent(e)
		}
	}()

	return r, nil
}

func (r *Remote) Publish(e event.Event) error {
	if r.conn == nil {
		return errors.New("connection not established")
	}

	b, err := websocketevent.ToJSON(e)
	if err != nil {
		return err
	}

	err = r.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		return err
	}

	return nil
}

func (r *Remote) handleEvent(e event.Event) {
	r.events <- e
}

func (r *Remote) Events() chan event.Event {
	return r.events
}
