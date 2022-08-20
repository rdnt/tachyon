package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"tachyon2/cmd/consume/event"
)

func main() {

}

type Event interface {
	Type() event.Type
}

func publishEvent(topic string, e any) error {
	b, err := json.Marshal(evt{
		Topic: topic,
		Type:  e,
		Event: nil,
	})
	if err != nil {
		return err
	}

	// return app.exchange.Publish(b)
}

type evt struct {
	Topic string
	Type  string
	Event []byte
}

func ConsumeEvent(b []byte) error {
	var e evt
	err := json.Unmarshal(b, &e)
	if err != nil {
		return err
	}

	switch e.Type {
	case event.SessionJoined:
		err = HandleSessionJoinedEvent(e.Event)
	case event.SessionLeft:
		err = HandleSessionLeftEvent(e.Event)
	case event.PathCreated:
		err = HandlePathCreatedEvent(e.Event)
	case event.PathTraced:
		err = HandlePathTracedEvent(e.Event)
	default:
		err = errors.New("invalid event type")
	}

	if err != nil {
		return err
	}

	return nil
}

func HandleSessionJoinedEvent(b []byte) error {
	var e event.SessionJoinedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return err
	}

	fmt.Println(e)
	return nil
}

func HandleSessionLeftEvent(b []byte) error {
	var e event.SessionLeftEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return err
	}

	fmt.Println(e)
	return nil
}

func HandlePathCreatedEvent(b []byte) error {
	var e event.PathCreatedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return err
	}

	fmt.Println(e)
	return nil
}

func HandlePathTracedEvent(b []byte) error {
	var e event.PathTracedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return err
	}

	fmt.Println(e)
	return nil
}
