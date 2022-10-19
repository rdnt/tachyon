package main_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"golang.org/x/exp/slices"
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

	coords := project.Vector2{
		X: 10,
		Y: 20,
	}

	t.Run("draw pixel", func(t *testing.T) {
		color, err := project.ColorFromString("#ff0000")
		assert.NilError(t, err)

		err = s.commands.DrawPixel(command.DrawPixelArgs{
			UserId:    uid,
			ProjectId: pid,
			Color:     color,
			Coords:    coords,
		})
		assert.NilError(t, err)

		eventually(t, func() bool {
			proj, err := s.projectRepo.Project(pid)
			if errors.Is(err, command.ErrProjectNotFound) {
				return false
			}
			assert.NilError(t, err)

			return slices.Contains(proj.Pixels, project.Pixel{
				Color:  color,
				Coords: coords,
			})
		})
	})

	t.Run("erase pixel", func(t *testing.T) {
		err := s.commands.ErasePixel(command.ErasePixelArgs{
			UserId:    uid,
			ProjectId: pid,
			Coords:    coords,
		})
		assert.NilError(t, err)

		eventually(t, func() bool {
			proj, err := s.projectRepo.Project(pid)
			if errors.Is(err, command.ErrProjectNotFound) {
				return false
			}
			assert.NilError(t, err)

			return len(proj.Pixels) == 0
		})
	})
}
