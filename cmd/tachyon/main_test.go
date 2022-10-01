package main_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/repository/project_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/session_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/user_repository"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/event_bus"
	"github.com/rdnt/tachyon/internal/event_store"
	"github.com/rdnt/tachyon/pkg/fanout"
	"gotest.tools/assert"
)

func TestTachyon(t *testing.T) {
	eventBus := event_bus.New(fanout.New[event.Event]())

	eventStore := event_store.New()

	sessionRepo, err := session_repository.New(eventStore)
	assert.NilError(t, err)

	userRepo, err := user_repository.New(eventStore)
	assert.NilError(t, err)

	projectRepo, err := project_repository.New(eventStore)
	assert.NilError(t, err)

	commandSvc := command.New(
		eventStore,
		eventBus,
		sessionRepo,
		projectRepo,
		userRepo,
	)

	//sessionView := session_view.New()
	//userView := user_view.New()
	//
	//querySvc := query.New(
	//	eventBus,
	//	sessionView,
	//	userView,
	//)

	uid := user.Id(uuid.New())
	t.Run("create user", func(t *testing.T) {
		name := "test user"
		err := commandSvc.CreateUser(uid, name)
		assert.NilError(t, err)

		u, err := userRepo.User(uid)
		assert.NilError(t, err)
		assert.Equal(t, u.Id, uid)
		assert.Equal(t, u.Name, name)

		t.Run("can't create user with the same name", func(t *testing.T) {
			err := commandSvc.CreateUser(uid, name)
			assert.Assert(t, err != nil)
		})
	})

	pid := project.Id(uuid.New())
	t.Run("create project", func(t *testing.T) {
		name := "first project"
		err := commandSvc.CreateProject(pid, name, uid)
		assert.NilError(t, err)

		p, err := projectRepo.Project(pid)
		assert.NilError(t, err)
		assert.Equal(t, p.Id, pid)
		assert.Equal(t, p.Name, name)
		assert.Equal(t, p.OwnerId, uid)

		p, err = projectRepo.UserProjectByName(uid, name)
		assert.NilError(t, err)
		assert.Equal(t, p.Id, pid)
		assert.Equal(t, p.Name, name)
		assert.Equal(t, p.OwnerId, uid)

		t.Run("this user can't create project with the same name", func(t *testing.T) {
			pid := project.Id(uuid.New())
			err := commandSvc.CreateProject(pid, name, uid)
			assert.Assert(t, err != nil)
		})
	})

	sid := session.Id(uuid.New())
	t.Run("create session", func(t *testing.T) {
		name := "first session"
		err := commandSvc.CreateSession(sid, name, pid)
		assert.NilError(t, err)

		s, err := sessionRepo.Session(sid)
		assert.NilError(t, err)
		assert.Equal(t, s.Id, sid)
		assert.Equal(t, s.Name, name)
		assert.Equal(t, s.ProjectId, pid)

		t.Run("project can't have session with the same name", func(t *testing.T) {
			sid := session.Id(uuid.New())
			err := commandSvc.CreateSession(sid, name, pid)
			assert.Assert(t, err != nil)
		})
	})
}
