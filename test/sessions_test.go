package test

import (
	"testing"

	"tachyon2/src/application"
	"tachyon2/src/application/domain/user"
	"tachyon2/src/application/event"
	"tachyon2/src/application/repositories/projectrepo"
	"tachyon2/src/application/repositories/sessionrepo"
	"tachyon2/src/application/repositories/userrepo"

	"gotest.tools/v3/assert"
)

func TestApplication(t *testing.T) {
	users := userrepo.New()
	projects := projectrepo.New()
	sessions := sessionrepo.New()

	app := application.New(users, projects, sessions)

	u1, err := users.CreateUser(user.New())
	assert.NilError(t, err)

	u2, err := users.CreateUser(user.New())
	assert.NilError(t, err)

	var joined, left int
	{
		dispose1 := app.Events().UserJoinedSession.Subscribe(func(e event.UserJoinedSession) {
			joined++
		})
		defer dispose1()

		dispose2 := app.Events().UserLeftSession.Subscribe(func(e event.UserLeftSession) {
			left++
		})
		defer dispose2()
	}

	s, err := app.CreateSession(u1.Id, "", "session")
	assert.NilError(t, err)

	{
		assert.Equal(t, joined, 0)

		err = app.JoinSession(u2.Id, s.Id)
		assert.NilError(t, err)

		assert.Equal(t, joined, 1)
	}

	{
		assert.Equal(t, left, 0)

		err = app.LeaveSession(u1.Id, s.Id)
		assert.NilError(t, err)

		assert.Equal(t, left, 1)
	}

	t.Log(s)
}
