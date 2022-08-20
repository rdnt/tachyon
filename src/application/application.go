package application

import (
	"encoding/json"
	"errors"
	"fmt"

	"tachyon2/pkg/fanout"
	"tachyon2/src/application/event"
	"tachyon2/src/application/repositories/projectrepo"
	"tachyon2/src/application/repositories/sessionrepo"
	"tachyon2/src/application/repositories/userrepo"
)

type AppEvent struct {
	AggregateId string
}

type App struct {
	id string

	users    *userrepo.Repository
	projects *projectrepo.Repository
	sessions *sessionrepo.Repository

	exchange fanout.Exchange[AppEvent]

	events Events
}

func New(id string, users *userrepo.Repository, projects *projectrepo.Repository, sessions *sessionrepo.Repository, exchange fanout.Exchange[AppEvent]) App {

	go func() {
		sub := exchange.Subscribe()

		for {
			select {
			case e := <-sub:
				fmt.Println("INCOMING EVENT FROM EXCHANGE")
				fmt.Println(e)
			}
		}
	}()

	return App{
		id: id,

		users:    users,
		projects: projects,
		sessions: sessions,

		exchange: exchange,
		events:   newEvents(id),
	}
}

type evt struct {
	Topic string
	Type  string
	Event []byte
}

// func (app *App) publishEvent(topic string, e any) error {
// 	evt := event{
// 		Topic: topic,
// 		Type:  e,
// 		Event: nil,
// 	}
//
// 	b, err := json.Marshal(evt)
// 	if err != nil {
// 		return err
// 	}
//
// 	return app.exchange.Publish(b)
// }

func (app *App) ConsumeEvent(b []byte) error {
	var e evt
	err := json.Unmarshal(b, &e)
	if err != nil {
		return err
	}

	switch e.Type {
	case event.SessionJoined:
		var e event.SessionJoinedEvent
		err := json.Unmarshal(b, &e)
		if err != nil {
			return err
		}

		app.events.SessionJoined.publish(e)
	case event.SessionLeft:
		var e event.SessionLeftEvent
		err := json.Unmarshal(b, &e)
		if err != nil {
			return err
		}

		app.events.SessionLeft.publish(e)
	case event.PathCreated:
		var e event.PathCreatedEvent
		err := json.Unmarshal(b, &e)
		if err != nil {
			return err
		}

		app.events.PathCreated.publish(e)
	case event.PathTraced:
		var e event.PathTracedEvent
		err := json.Unmarshal(b, &e)
		if err != nil {
			return err
		}

		app.events.PathTraced.publish(e)
	default:
		err = errors.New("invalid event type")
	}

	if err != nil {
		return err
	}

	return nil
}

// func ConsumeEvent(b []byte) error {
// 	var e evt
// 	err := json.Unmarshal(b, &e)
// 	if err != nil {
// 		return err
// 	}
//
// 	switch e.Type {
// 	case event.SessionJoined:
// 		err = HandleSessionJoinedEvent(e.Event)
// 	case event.SessionLeft:
// 		err = HandleSessionLeftEvent(e.Event)
// 	case event.PathCreated:
// 		err = HandlePathCreatedEvent(e.Event)
// 	case event.PathTraced:
// 		err = HandlePathTracedEvent(e.Event)
// 	default:
// 		err = errors.New("invalid event type")
// 	}
//
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

// func (app App) Events() Events {
// 	return app.event
// }
