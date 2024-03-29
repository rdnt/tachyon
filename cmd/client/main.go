package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"tachyon/internal/client/application"
	"tachyon/internal/client/remote"
	"tachyon/internal/client/repository/project_repository"
	"tachyon/internal/client/repository/session_repository"
	"tachyon/internal/client/repository/user_repository"
)

func main() {
	r, err := remote.New("ws://0.0.0.0:8080/ws")
	if err != nil {
		panic(err)
	}

	sessionRepo, err := session_repository.New()
	if err != nil {
		log.Fatal(err)
	}

	userRepo, err := user_repository.New()
	if err != nil {
		log.Fatal(err)
	}

	projectRepo, err := project_repository.New()
	if err != nil {
		log.Fatal(err)
	}

	app, err := application.New(r, sessionRepo, projectRepo, userRepo)
	if err != nil {
		panic(err)
	}

	// err = app.CreateUser("user-1")
	// fmt.Println(err)
	//
	// err = app.CreateProject("project-1")
	// fmt.Println(err)

	m := &model{
		app: app,
		//userId:    uuid.Nil,
		//projectId: uuid.Nil, // TODO: no need
	}

	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseAllMotion())

	//_ = p
	//stop := make(chan os.Signal, 1)
	//signal.Notify(stop, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGHUP)

	//{
	//	err = m.app.CreateSession("my-session")
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	time.Sleep(1 * time.Second)
	//
	//	sess := m.app.Session()
	//
	//	err = m.app.CreatePath(sess.ProjectId,
	//		project.Color{
	//			R: 0xff, G: 0xff, B: 0xff, A: 0xff,
	//		},
	//		project.Vector2{
	//			X: float64(3),
	//			Y: float64(3),
	//		})
	//
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//}

	_, err = p.Run()
	if err != nil {
		log.Fatal(err)
	}

	//}()
	//<-stop

}
