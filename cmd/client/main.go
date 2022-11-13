package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rdnt/tachyon/internal/client/application"
	"github.com/rdnt/tachyon/internal/client/remote"
	"github.com/rdnt/tachyon/pkg/uuid"
)

func main() {
	r, err := remote.New("localhost:1967")
	if err != nil {
		panic(err)
	}

	app, err := application.New(r)
	if err != nil {
		panic(err)
	}

	err = app.CreateUser("user-1")
	fmt.Println(err)

	err = app.CreateProject("project-1")
	fmt.Println(err)

	m := &model{
		app:       app,
		projectId: uuid.Nil, // TODO: no need
	}

	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseAllMotion())

	err = p.Start()
	if err != nil {
		log.Fatal(err)
	}
}
