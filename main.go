package main

import (
	"fmt"
	"os"
	"os/signal"

	"tachyon2/src/application"
	"tachyon2/src/application/domain/user"
	"tachyon2/src/application/event"
	"tachyon2/src/application/repositories/projectrepo"
	"tachyon2/src/application/repositories/sessionrepo"
	"tachyon2/src/application/repositories/userrepo"
)

func main() {
	users := userrepo.New()
	projects := projectrepo.New()
	sessions := sessionrepo.New()

	app := application.New(users, projects, sessions)

	u, err := users.CreateUser(user.User{
		Id:      "1",
		Name:    "u",
		Pointer: user.Pointer{},
	})
	if err != nil {
		panic(err)
	}

	u2, err := users.CreateUser(user.User{
		Id:      "2",
		Name:    "u2",
		Pointer: user.Pointer{},
	})
	if err != nil {
		panic(err)
	}

	sess, err := app.CreateSession(u.Id, "test", "my-session")
	if err != nil {
		panic(err)
	}

	app.Events().UserJoinedSession.Subscribe(func(e event.UserJoinedSession) {
		fmt.Println("User joined!!!! :D")
	})

	err = app.JoinSession(u2.Id, sess.Id)
	if err != nil {
		panic(err)
	}

	sess, _ = app.Session(sess.Id)
	fmt.Println(sess)

	err = app.LeaveSession(u2.Id, sess.Id)
	if err != nil {
		panic(err)
	}

	sess, _ = app.Session(sess.Id)
	fmt.Println(sess)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
