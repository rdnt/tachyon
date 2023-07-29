package path

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"tachyon/pkg/uuid"
)

type Path struct {
	Id     uuid.UUID
	Tool   Tool
	Color  Color
	Points []Vector2
}

type Vector2 struct {
	X float64
	Y float64
}

type Tool string

const (
	Pen    Tool = "pen"
	Eraser Tool = "eraser"
)

func (t Tool) String() string {
	return string(t)
}

type Color color.RGBA

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

func New(id uuid.UUID, tool Tool, color Color, points []Vector2) Path {
	return Path{
		Id:     id,
		Tool:   tool,
		Color:  color,
		Points: points,
	}
}
