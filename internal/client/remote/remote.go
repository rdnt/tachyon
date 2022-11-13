package remote

import (
	"errors"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/rdnt/tachyon/internal/client/application/event"
	"github.com/rdnt/tachyon/internal/client/remote/websocketevent"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type Remote struct {
	address  string
	conn     *websocket.Conn
	messages chan event.Event
}

func New(address string) (*Remote, error) {
	r := &Remote{
		address:  address,
		messages: make(chan event.Event),
	}

	go func() {
		for {
			func() {
				conn, _, err := websocket.DefaultDialer.Dial(address, nil)
				if err != nil {
					log.Fatal("dial:", err)
				}
				defer conn.Close()

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

	return r.conn.WriteMessage(websocket.TextMessage, b)
}

func (r *Remote) handleEvent(e event.Event) {
	switch e.Type() {
	case event.UserCreated:
		fmt.Println("USER CREATED", e)
	}

	fmt.Println("received event", e)
}

func (r *Remote) Project(id uuid.UUID) error {
	return nil
}
