package test

import (
	"testing"

	"github.com/rdnt/tachyon/test/suite"
	"gotest.tools/assert"
)

func TestIntegration(t *testing.T) {
	s := suite.New(t)

	err := s.Client.App.CreateUser("user-1")
	assert.NilError(t, err)

	err = s.Client.App.CreateProject("project-1")
	assert.NilError(t, err)
}
