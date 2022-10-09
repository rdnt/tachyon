package main_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"golang.org/x/exp/slices"
	"gotest.tools/assert"
)

func TestSessions(t *testing.T) {
	s := newSuite(t)

	uid1 := user.Id(uuid.New())
	uid2 := user.Id(uuid.New())

	t.Run("create users", func(t *testing.T) {
		err := s.commands.CreateUser(uid1, "user-1")
		assert.NilError(t, err)

		err = s.commands.CreateUser(uid2, "user-2")
		assert.NilError(t, err)
	})

	pid := project.Id(uuid.New())

	t.Run("create project", func(t *testing.T) {
		err := s.commands.CreateProject(pid, "project-1", uid1)
		assert.NilError(t, err)
	})

	sid := session.Id(uuid.New())

	t.Run("create session", func(t *testing.T) {
		err := s.commands.CreateSession(sid, "session-1", pid)
		assert.NilError(t, err)

		s, err := s.sessionRepo.Session(sid)
		assert.NilError(t, err)

		assert.Assert(t, slices.Contains(s.UserIds, uid1))
		assert.Equal(t, len(s.UserIds), 1)
	})

	t.Run("user2 joins session", func(t *testing.T) {
		err := s.commands.JoinSession(sid, uid2)
		assert.NilError(t, err)

		s, err := s.sessionRepo.Session(sid)
		assert.NilError(t, err)

		assert.Assert(t, slices.Contains(s.UserIds, uid1))
		assert.Assert(t, slices.Contains(s.UserIds, uid2))
		assert.Equal(t, len(s.UserIds), 2)

		// TODO: verify queries received update
	})

	t.Run("user2 leaves session", func(t *testing.T) {
		err := s.commands.LeaveSession(sid, uid2)
		assert.NilError(t, err)

		s, err := s.sessionRepo.Session(sid)
		assert.NilError(t, err)

		assert.Assert(t, slices.Contains(s.UserIds, uid1))
		assert.Assert(t, !slices.Contains(s.UserIds, uid2))
		assert.Equal(t, len(s.UserIds), 1)

		// TODO: verify queries received update
	})
}
