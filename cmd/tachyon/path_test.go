package main_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/project/path"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"gotest.tools/assert"
)

func TestPath(t *testing.T) {
	s := newSuite(t)

	uid := user.Id(uuid.New())
	t.Run("create user", func(t *testing.T) {
		err := s.commands.CreateUser(uid, "user-1")
		assert.NilError(t, err)
	})

	pid := project.Id(uuid.New())
	t.Run("create project", func(t *testing.T) {
		err := s.commands.CreateProject(pid, "project-1", uid)
		assert.NilError(t, err)
	})

	pathId := path.Id(uuid.New())
	t.Run("create path", func(t *testing.T) {
		tool := path.Pen

		color, err := path.ColorFromString("#ff0000")
		assert.NilError(t, err)

		point := path.Vector2{
			X: 10,
			Y: 20.03,
		}

		err = s.commands.CreatePath(command.CreatePathArgs{
			PathId:    pathId,
			UserId:    uid,
			ProjectId: pid,
			Tool:      tool,
			Color:     color,
			Point:     point,
		})
		assert.NilError(t, err)

		eventually(t, func() bool {
			return true
		})
	})
}
