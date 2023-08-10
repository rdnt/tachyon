package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	gookitcolor "github.com/gookit/color"

	"tachyon/internal/client/application"
	"tachyon/internal/client/application/domain/project"
)

type model struct {
	width  int
	height int
	app    *application.Application
	debug  string
}

func (m *model) Init() tea.Cmd {
	return tea.Batch(
		tick(),
	)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		//		cmds = append(cmds, resizeCmd)
	case tea.MouseMsg:
		switch {
		case msg.Action == tea.MouseActionMotion && msg.Button == tea.MouseButtonLeft:
			m.debug = fmt.Sprintf("%d   %d   %d   %d", m.width, m.height, msg.X, msg.Y)

			sess, err := m.app.Session()
			if err == nil {
				err := m.app.CreatePath(
					sess.ProjectId,
					project.Color{
						R: 0xff, G: 0xff, B: 0xff, A: 0xff,
					},
					project.Vector2{
						X: float64(msg.X),
						Y: float64(msg.Y),
					},
				)
				if err != nil {
					panic(err)
				}
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "c":
			err := m.app.CreateSession("my-session")
			if err != nil {
				fmt.Println("@@@@ errrr", err)
			}
			// case "p":
			// 	err := m.app.CreateProject("my-project")
			// 	if err != nil {
			// 		panic(err)
			// 	}
		}

	case time.Time:
		cmds = append(cmds, tick())
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	proj, err := m.app.Project()
	if err != nil {
		return ""
	}

	canvas := make([][]gookitcolor.RGBColor, m.height)
	for i := range canvas {
		canvas[i] = make([]gookitcolor.RGBColor, m.width)
	}

	for _, p := range proj.Paths {
		if p.Points[0].Y >= float64(m.height) || p.Points[0].X >= float64(m.width) || p.Points[0].X < 0 || p.Points[0].Y < 0 {
			continue
		}

		canvas[int(p.Points[0].Y)][int(p.Points[0].X)] = gookitcolor.Hex(fmt.Sprintf("#%02x%02x%02x", p.Color.R, p.Color.G, p.Color.B), true)
	}

	var s strings.Builder
	for i, row := range canvas {
		for j, clr := range row {
			if i == 0 && j < len(m.debug) {
				s.WriteString(string(rune(m.debug[j])))
			} else {
				s.WriteString(clr.Sprint(" "))
			}
		}
		s.WriteString("\n")
	}

	return strings.TrimSuffix(s.String(), "\n")
}

func tick() tea.Cmd {
	return tea.Tick(
		16*time.Millisecond, func(t time.Time) tea.Msg {
			return t
		},
	)
}

// var screen, _ = tcell.NewScreen()
//
// func init() {
//	if err := screen.Init(); err != nil {
//		return
//	}
// }
//
// var resizeCmd = tea.Tick(time.Second/33, func(time.Time) tea.Msg {
//	//	screen, _ := tcell.NewScreen()
//
//	//	defer screen.Fini()
//	w, h := screen.Size()
//	return tea.WindowSizeMsg{
//		Width:  w,
//		Height: h,
//	}
// })
