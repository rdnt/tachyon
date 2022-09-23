package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/rdnt/tachyon/internal/application"
	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/repository/session_repository"
	"github.com/rdnt/tachyon/internal/application/domain/session"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/application/query"
	readSessionRepo "github.com/rdnt/tachyon/internal/application/query/view/session_view"
	"github.com/rdnt/tachyon/internal/event_bus"
	"github.com/rdnt/tachyon/internal/event_store"
	"github.com/rdnt/tachyon/pkg/fanout"
)

func main() {
	time.Sleep(1 * time.Second)
	eventStore := event_store.New()
	sessionRepository := session_repository.New(eventStore)
	eventBus := event_bus.New(fanout.New[event.Event]())

	cmds := command.New(eventStore, eventBus, sessionRepository)

	//rsr := readSessionRepo.New(eventBus)
	//qs := query.New(rsr)

	qs := []application.Queries{}
	for i := 0; i < 10; i++ {
		rsr := readSessionRepo.New(eventBus)
		qs = append(qs, query.New(rsr))
	}

	//app := application.New(cmds, qs)

	err := cmds.CreateSession("userId", "projectId", "sessionName")
	if err != nil {
		panic(err)
	}

	//log.Debugln("eventStore", eventStore)
	//log.Debugln("sessionRepository", sessionRepository)

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	sessionId := session.Id(input.Text())

	for i := 0; i < 10; i++ {
		fmt.Println(qs[i].Session(sessionId))
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
