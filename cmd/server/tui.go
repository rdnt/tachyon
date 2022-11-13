package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	gookitcolor "github.com/gookit/color"
	"github.com/rdnt/tachyon/internal/server/application/command"
	"github.com/rdnt/tachyon/internal/server/application/domain/project"
	"github.com/rdnt/tachyon/internal/server/application/query"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type model struct {
	width     int
	height    int
	commands  command.Service
	queries   query.Service
	userId    uuid.UUID
	projectId uuid.UUID
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
	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseLeft:
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

	case time.Time:
		cmds = append(cmds, tick())
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	proj, err := m.queries.Project(m.projectId)
	if err != nil {
		return ""
	}

	canvas := make([][]gookitcolor.RGBColor, m.height)
	for i := range canvas {
		canvas[i] = make([]gookitcolor.RGBColor, m.width)
	}

	for _, p := range proj.Pixels {
		if p.Coords.Y >= m.height || p.Coords.X >= m.width || p.Coords.Y < 0 || p.Coords.X < 0 {
			continue
		}

		canvas[p.Coords.Y][p.Coords.X] = gookitcolor.Hex(fmt.Sprintf("#%02x%02x%02x", p.Color.R, p.Color.G, p.Color.B), true)
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

func tick() tea.Cmd {
	return tea.Tick(
		16*time.Millisecond, func(t time.Time) tea.Msg {
			return t
		},
	)
}
