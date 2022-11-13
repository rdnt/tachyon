package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestIntegration(t *testing.T) {
	newServer(t)
	c := newClient(t)

	err := c.app.CreateUser("user-1")
	assert.NilError(t, err)

	err = c.app.CreateProject("project-1")
	assert.NilError(t, err)
}
