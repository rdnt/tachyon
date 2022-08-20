package main

import (
	"fmt"
	"os"
	"os/signal"

	"tachyon2/pkg/broker"
	"tachyon2/pkg/fanout"
	"tachyon2/src/application"
	"tachyon2/src/application/domain/user"
	"tachyon2/src/application/repositories/projectrepo"
	"tachyon2/src/application/repositories/sessionrepo"
	"tachyon2/src/application/repositories/userrepo"
)

type UserJoined struct {
	J bool
}

type UserLeft struct {
	L string
}

type Event struct {
	Id      string
	Channel string
	Payload []byte
}

type EventBroker struct {
	channels map[string]struct {
		UserJoinedEvent *broker.Broker[UserJoined]
		UserLeftEvent   *broker.Broker[UserLeft]
	}
}

func main() {
	instances := 3
	apps := make([]application.App, instances)

	users := userrepo.New()
	projects := projectrepo.New()
	sessions := sessionrepo.New()

	exchange := fanout.New[application.AppEvent]()

	u0, err := users.CreateUser(user.New())
	handle(err)

	u1, err := users.CreateUser(user.New())
	handle(err)

	u2, err := users.CreateUser(user.New())
	handle(err)

	for i := 0; i < 3; i++ {
		apps[i] = application.New(fmt.Sprint(i), users, projects, sessions, exchange)
	}

	{
		sess, err := apps[0].CreateSession(u0.Id, "", "my-session")
		handle(err)

		fmt.Println(sess)

		sess1, err := apps[1].JoinSession(u1.Id, sess.Id)
		handle(err)

		fmt.Println(sess1)

		sess2, err := apps[2].JoinSession(u2.Id, sess.Id)
		handle(err)

		fmt.Println(sess2)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
