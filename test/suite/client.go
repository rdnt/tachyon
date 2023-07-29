package suite

import (
	"testing"

	"gotest.tools/assert"

	"tachyon/internal/client/application"
	"tachyon/internal/client/remote"
)

type Client struct {
	App *application.Application
}

func NewClient(t *testing.T) *Client {
	r, err := remote.New("ws://:80/ws")
	assert.NilError(t, err)

	app, err := application.New(r, nil, nil, nil)
	assert.NilError(t, err)

	return &Client{
		App: app,
	}
}

func (c *Client) Close() error {
	return nil
}
