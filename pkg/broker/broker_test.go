package broker

import (
	"testing"

	"gotest.tools/v3/assert"
)

type TestEvent struct {
	id string
}

func TestBroker(t *testing.T) {
	testId := "test-id"

	evt := TestEvent{
		id: testId,
	}

	broker := New[TestEvent]()

	var received bool
	broker.Subscribe(func(e TestEvent) {
		if e.id == testId {
			received = true
		}
	})

	broker.Publish(evt)

	assert.Equal(t, received, true)
}
