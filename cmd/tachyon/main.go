package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/google/uuid"
	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/repository/project_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/session_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/user_repository"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/query"
	"github.com/rdnt/tachyon/internal/application/query/view/session_view"
	"github.com/rdnt/tachyon/internal/application/query/view/user_view"
	"github.com/rdnt/tachyon/internal/event_bus"
	"github.com/rdnt/tachyon/internal/event_store"
	"github.com/rdnt/tachyon/pkg/fanout"
)

func main() {
	eventBus := event_bus.New(fanout.New[Event]())

	eventStore := event_store.New()
	sessionRepo, err := session_repository.New(eventStore)
	if err != nil {
		panic(err)
	}
	userRepo, err := user_repository.New(eventStore)
	if err != nil {
		panic(err)
	}
	projectRepo, err := project_repository.New(eventStore)
	if err != nil {
		panic(err)
	}

	commandSvc := command.New(
		eventStore,
		eventBus,
		sessionRepo,
		projectRepo,
		userRepo,
	)

	sessionView := session_view.New()
	userView := user_view.New()

	projectView, err := project_repository.New(eventStore)
	if err != nil {
		panic(err)
	}

	querySvc := query.New(
		eventBus,
		sessionView,
		userView,
		projectView,
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
				uid := uuid.New()
				err := commandSvc.CreateUser(user.Id(uid), "user-1")
				if err != nil {
					panic(err)
				}

				ru, err := userRepo.User(user.Id(uid))
				if err != nil {
					panic(err)
				}
				fmt.Println("USER FOUND", ru)

				//err = userRepo.Hydrate()
				//if err != nil {
				//	panic(err)
				//}

				//for i := 0; i < 1000; i++ {
				//	u, err := querySvc.User("user-1")
				//	if err != nil {
				//		fmt.Println(err)
				//		continue
				//	}
				//
				//	fmt.Println(u)
				//	break
				//}

			case "create project":
				pid := uuid.New()

				fmt.Println("name?")
				input.Scan()
				name := input.Text()

				fmt.Println("ownerId?")
				input.Scan()
				uid := uuid.MustParse(input.Text())

				err := commandSvc.CreateProject(project.Id(pid), name, user.Id(uid))
				if err != nil {
					panic(err)
				}

			//case "create session":
			//	err := commandSvc.CreateSession("someUserId", "someProjectId", "my-session")
			//	if err != nil {
			//		panic(err)
			//	}
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
