package application_test

import (
	"errors"
	"testing"

	"tachyon/internal/server/application/query/view/session_view"
	"tachyon/pkg/uuid"
	"golang.org/x/exp/slices"
	"gotest.tools/assert"
)

func TestSessions(t *testing.T) {
	s := newSuite(t)

	uid1 := uuid.New()
	uid2 := uuid.New()

	t.Run("create users", func(t *testing.T) {
		err := s.commands.CreateUser(uid1, "user-1")
		assert.NilError(t, err)

		err = s.commands.CreateUser(uid2, "user-2")
		assert.NilError(t, err)
	})

	pid := uuid.New()

	t.Run("create project", func(t *testing.T) {
		err := s.commands.CreateProject(pid, "project-1", uid1)
		assert.NilError(t, err)
	})

	sid := uuid.New()

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

		sess, err := s.sessionRepo.Session(sid)
		assert.NilError(t, err)

		assert.Assert(t, slices.Contains(sess.UserIds, uid1))
		assert.Assert(t, slices.Contains(sess.UserIds, uid2))
		assert.Equal(t, len(sess.UserIds), 2)

		eventually(t, func() bool {
			sess, err := s.queries.Session(sid)
			if errors.Is(err, session_view.ErrSessionNotFound) {
				return false
			}
			assert.NilError(t, err)

			if len(sess.UserIds) != 2 {
				return false
			}

			if !slices.Contains(sess.UserIds, uid1) {
				return false
			}

			if !slices.Contains(sess.UserIds, uid2) {
				return false
			}

			return true
		})
	})

	t.Run("user2 leaves session", func(t *testing.T) {
		err := s.commands.LeaveSession(sid, uid2)
		assert.NilError(t, err)

		sess, err := s.sessionRepo.Session(sid)
		assert.NilError(t, err)

		assert.Assert(t, slices.Contains(sess.UserIds, uid1))
		assert.Assert(t, !slices.Contains(sess.UserIds, uid2))
		assert.Equal(t, len(sess.UserIds), 1)

		eventually(t, func() bool {
			sess, err := s.queries.Session(sid)
			if errors.Is(err, session_view.ErrSessionNotFound) {
				return false
			}
			assert.NilError(t, err)

			if len(sess.UserIds) != 1 {
				return false
			}

			if !slices.Contains(sess.UserIds, uid1) {
				return false
			}

			return true
		})
	})
}
