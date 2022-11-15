package suite

import (
	"testing"

	"gotest.tools/assert"
)

type Suite struct {
	Server *Server
	Client *Client
}

func New(t *testing.T) *Suite {
	s := &Suite{
		Server: NewServer(t),
		Client: NewClient(t),
	}

	t.Cleanup(func() {
		err := s.Client.Close()
		assert.NilError(t, err)

		err = s.Server.Close()
		assert.NilError(t, err)
	})

	return s
}
