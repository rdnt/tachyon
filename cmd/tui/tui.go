package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	gookitcolor "github.com/gookit/color"
	"github.com/rdnt/tachyon/internal/application/command"
	"github.com/rdnt/tachyon/internal/application/domain/project"
	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/internal/application/query"
)

type model struct {
	width     int
	height    int
	commands  command.Service
	queries   query.Service
	userId    user.Id
	projectId project.Id
}

// Init initializes the model with the initial state.
func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// On window resize, we want to create a replica grid
		// (called the canvas) for the user to draw on.
		//
		// This will be a 2D slice of strings. We use strings and not runes so
		// that we can store the style of the character drawn as well so that
		// each cell can be a different style / color.
		m.width = msg.Width
		m.height = msg.Height
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseLeft:
			// When the user clicks on the mouse, we want to write the
			// character to the current position of the mouse in the grid, so
			// that we can draw it later.
			//m.canvas[msg.Y][msg.X] = gookitcolor.Hex("#ffffff", true)
			err := m.commands.DrawPixel(command.DrawPixelArgs{
				UserId:    m.userId,
				ProjectId: m.projectId,
				Color: project.Color{
					R: 0xff, G: 0xff, B: 0xff, A: 0xff,
				},
				Coords: project.Vector2{
					X: msg.X,
					Y: msg.Y,
				},
			})
			if err != nil {
				panic(err)
			}
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	canvas := make([][]gookitcolor.RGBColor, m.height)
	for i := range canvas {
		canvas[i] = make([]gookitcolor.RGBColor, m.width)
	}

	proj, err := m.queries.Project(m.projectId)
	if err != nil {
		panic(err)
	}

	for _, p := range proj.Pixels {
		if p.Coords.Y >= len(canvas) || p.Coords.X >= len(canvas[p.Coords.Y]) {
			canvas[p.Coords.Y][p.Coords.X] = gookitcolor.Hex(fmt.Sprintf("#%02x%02x%02x", p.Color.R, p.Color.G, p.Color.B), true)
		}
	}

	var s strings.Builder
	for _, row := range canvas {
		for _, clr := range row {
			s.WriteString(clr.Sprint(" "))
		}
		s.WriteString("\n")
	}
	return strings.TrimSuffix(s.String(), "\n")

}
