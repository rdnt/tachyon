package main_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"gotest.tools/assert"
)

func TestDrawPixel(t *testing.T) {
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

	t.Run("draw pixel", func(t *testing.T) {
		color, err := project.ColorFromString("#ff0000")
		assert.NilError(t, err)

		coords := project.Vector2{
			X: 10,
			Y: 20,
		}

		err = s.commands.DrawPixel(command.DrawPixelArgs{
			UserId:    uid,
			ProjectId: pid,
			Color:     color,
			Coords:    coords,
		})
		assert.NilError(t, err)

		eventually(t, func() bool {
			return true
		})
	})
}

func TestErasePixel(t *testing.T) {
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

	t.Run("erase pixel", func(t *testing.T) {
		coords := project.Vector2{
			X: 10,
			Y: 20,
		}

		err := s.commands.ErasePixel(command.ErasePixelArgs{
			UserId:    uid,
			ProjectId: pid,
			Coords:    coords,
		})
		assert.NilError(t, err)

		eventually(t, func() bool {
			return true
		})
	})
}
