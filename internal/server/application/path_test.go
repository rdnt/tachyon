package application_test

import (
	"testing"

	"gotest.tools/assert"

	"tachyon/internal/server/application/command"
	"tachyon/internal/server/application/domain/project/path"
	"tachyon/pkg/uuid"
)

func TestPathTrace(t *testing.T) {
	s := newSuite(t)

	uid := uuid.New()
	t.Run("create user", func(t *testing.T) {
		err := s.commands.CreateUser(uid, "user-1")
		assert.NilError(t, err)
	})

	pid := uuid.New()
	t.Run("create project", func(t *testing.T) {
		err := s.commands.CreateProject(pid, "project-1", uid)
		assert.NilError(t, err)
	})

	pathId := uuid.New()
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

	t.Run("trace path", func(t *testing.T) {
		point := path.Vector2{
			X: 10,
			Y: 20.03,
		}

		err := s.commands.TracePath(command.TracePathArgs{
			PathId:    pathId,
			UserId:    uid,
			ProjectId: pid,
			Point:     point,
		})
		assert.NilError(t, err)

		eventually(t, func() bool {
			return true
		})
	})
}
