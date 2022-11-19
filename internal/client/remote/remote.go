package remote

import (
	"encoding/json"
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

type jsonEvent struct {
	Type string `json:"type"`
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

			fmt.Println("Unmarshal", string(b))

			var tmp jsonEvent
			err = json.Unmarshal(b, &tmp)
			if err != nil {
				fmt.Println("invalid message format1")
				continue
			}

			var e event.Event
			// TODO: @rdnt can't do that. add marshal unmarshal switch to shared pkg
			err = json.Unmarshal(b, &e)
			if err != nil {
				fmt.Println("invalid message format2", err)
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

	b, err := json.Marshal(e)
	if err != nil {
		return err
	}

	var tmp map[string]any
	err = json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	tmp["type"] = e.Type()

	b, err = json.Marshal(tmp)
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

func (r *Remote) HandleConnectedEvent(e event.ConnectedEvent) {
	fmt.Println("CONNECTED", e)
}
