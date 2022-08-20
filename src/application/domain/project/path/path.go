package path

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"tachyon2/pkg/uuid"
)

type Id string

type Path struct {
	Id     Id
	Tool   Tool
	Color  Color
	Points []Point
}

type Tool string

const (
	Pen    Tool = "pen"
	Eraser Tool = "eraser"
)

type Color color.RGBA

type Point struct {
	X float64
	Y float64
}

func (c Color) String() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

func ColorFromString(s string) (c Color, err error) {
	values, err := strconv.ParseUint(strings.TrimLeft(s, "#"), 16, 32)
	if err != nil {
		return Color{}, err
	}

	return Color{
		R: uint8(values >> 16),
		G: uint8((values >> 8) & 0xFF),
		B: uint8(values & 0xFF),
		A: uint8(0xFF),
	}, nil
}

func (p *Path) AddPoint(x, y float64) {
	p.Points = append(p.Points, Point{
		X: x,
		Y: y,
	})
}

func New(tool Tool, color Color, x, y float64) Path {
	return Path{
		Id:    Id(uuid.New()),
		Tool:  tool,
		Color: color,
		Points: []Point{
			{
				X: x,
				Y: y,
			},
		},
	}
}
