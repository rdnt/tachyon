package main_test

import (
	"errors"
	"testing"

	"github.com/rdnt/tachyon/internal/server/application/query/view/session_view"
	"github.com/rdnt/tachyon/pkg/uuid"
	"gotest.tools/assert"
)

func TestTachyon(t *testing.T) {
	s := newSuite(t)

	uid := uuid.New()
	t.Run("create user", func(t *testing.T) {
		name := "test user"
		err := s.commands.CreateUser(uid, name)
		assert.NilError(t, err)

		u, err := s.userRepo.User(uid)
		assert.NilError(t, err)
		assert.Equal(t, u.Id, uid)
		assert.Equal(t, u.Name, name)

		t.Run("can't create user with the same name", func(t *testing.T) {
			err := s.commands.CreateUser(uid, name)
			assert.Assert(t, err != nil)
		})
	})

	pid := uuid.New()
	t.Run("create project", func(t *testing.T) {
		name := "first project"
		err := s.commands.CreateProject(pid, name, uid)
		assert.NilError(t, err)

		p, err := s.projectRepo.Project(pid)
		assert.NilError(t, err)
		assert.Equal(t, p.Id, pid)
		assert.Equal(t, p.Name, name)
		assert.Equal(t, p.OwnerId, uid)

		p, err = s.projectRepo.UserProjectByName(uid, name)
		assert.NilError(t, err)
		assert.Equal(t, p.Id, pid)
		assert.Equal(t, p.Name, name)
		assert.Equal(t, p.OwnerId, uid)

		t.Run("this user can't create project with the same name", func(t *testing.T) {
			pid := uuid.New()
			err := s.commands.CreateProject(pid, name, uid)
			assert.Assert(t, err != nil)
		})
	})

	sid := uuid.New()
	t.Run("create session", func(t *testing.T) {
		name := "first session"
		err := s.commands.CreateSession(sid, name, pid)
		assert.NilError(t, err)

		sess, err := s.sessionRepo.Session(sid)
		assert.NilError(t, err)
		assert.Equal(t, sess.Id, sid)
		assert.Equal(t, sess.Name, name)
		assert.Equal(t, sess.ProjectId, pid)

		t.Run("session can be queried", func(t *testing.T) {
			eventually(t, func() bool {
				sess, err := s.queries.Session(sid)
				if errors.Is(err, session_view.ErrSessionNotFound) {
					return false
				}
				assert.NilError(t, err)

				return sess.Id == sid
			})
		})

		t.Run("project can't have session with the same name", func(t *testing.T) {
			sid := uuid.New()
			err := s.commands.CreateSession(sid, name, pid)
			assert.Assert(t, err != nil)
		})
	})
}
