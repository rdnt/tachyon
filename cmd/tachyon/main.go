package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/repository/session_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/user_repository"
	"github.com/rdnt/tachyon/internal/application/event"
	"github.com/rdnt/tachyon/internal/application/query"
	"github.com/rdnt/tachyon/internal/application/query/view/session_view"
	"github.com/rdnt/tachyon/internal/application/query/view/user_view"
	"github.com/rdnt/tachyon/internal/event_bus"
	"github.com/rdnt/tachyon/internal/event_store"
	"github.com/rdnt/tachyon/pkg/fanout"
)

func main() {
	eventBus := event_bus.New(fanout.New[event.Event]())

	eventStore := event_store.New()
	sessionRepo := session_repository.New(eventStore)
	userRepo := user_repository.New(eventStore)

	commandSvc := command.New(
		eventStore,
		eventBus,
		sessionRepo,
		userRepo,
	)

	sessionView := session_view.New()
	userView := user_view.New()

	querySvc := query.New(
		eventBus,
		sessionView,
		userView,
	)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		_ = commandSvc
		_ = querySvc

		for {
			fmt.Print("> ")
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			cmd := input.Text()

			fmt.Println()

			switch cmd {
			case "store":
				fmt.Println(eventStore)
			case "bus":
				fmt.Println(eventBus)
			case "repo":
				fmt.Println("sessions:", sessionRepo)
				fmt.Println("users:", userRepo)
			case "view":
				fmt.Println("sessions:", sessionView)
				fmt.Println("users:", userView)
			case "create user":
				err := commandSvc.CreateUser("tasos")
				if err != nil {
					panic(err)
				}

				ru, err := userRepo.User("tasos")
				if err != nil {
					panic(err)
				}

				fmt.Println("USER FOUND", ru)

				//for i := 0; i < 1000; i++ {
				//	u, err := querySvc.User("tasos")
				//	if err != nil {
				//		fmt.Println(err)
				//		continue
				//	}
				//
				//	fmt.Println(u)
				//	break
				//}

			case "create session":
				err := commandSvc.CreateSession("someUserId", "someProjectId", "my-session")
				if err != nil {
					panic(err)
				}
			default:
				fmt.Println("invalid command")
			case "quit":
				close(c)
			}

			time.Sleep(1 * time.Millisecond)

			fmt.Println()
		}
	}()

	<-c
}
