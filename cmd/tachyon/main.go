package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/command/repository/project_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/session_repository"
	"github.com/rdnt/tachyon/internal/application/command/repository/user_repository"
	"github.com/rdnt/tachyon/internal/application/query"
	"github.com/rdnt/tachyon/internal/pkg/redis/client"
	"github.com/rdnt/tachyon/internal/pkg/redis/eventbus"
	"github.com/rdnt/tachyon/internal/pkg/redis/eventstore"
	"github.com/rdnt/tachyon/pkg/uuid"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
		DB:   0,
	})

	const redisStreamKey = "events"
	redisClient := client.New(rdb, redisStreamKey)
	eventStore := eventstore.New(redisClient)
	eventBus := eventbus.New(redisClient)
	//eventBus := event_bus.New(fanout.New[event.Event]())

	//eventStore := event_store.New()
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

	sessionView, err := session_repository.New(eventBus)
	if err != nil {
		panic(err)
	}

	userView, err := user_repository.New(eventBus)
	if err != nil {
		panic(err)
	}

	projectView, err := project_repository.New(eventBus)
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
				fmt.Println("projects:", projectRepo)
			case "view":
				fmt.Println("sessions:", sessionView)
				fmt.Println("users:", userView)
				fmt.Println("projects:", projectView)
			case "create user":
				uid := uuid.New()
				err := commandSvc.CreateUser(uid, "user-1")
				if err != nil {
					panic(err)
				}

				ru, err := userRepo.User(uid)
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
				uid, err := uuid.Parse(input.Text())
				if err != nil {
					panic(err)
				}

				err = commandSvc.CreateProject(pid, name, uid)
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
