package path

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"github.com/rdnt/tachyon/internal/application/domain/project/path/point"
)

type Id string

type Path struct {
	Id     Id
	Tool   Tool
	Color  Color
	Points []point.Id
}

type Tool string

const (
	Pen    Tool = "pen"
	Eraser Tool = "eraser"
)

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

func New(id Id, tool Tool, color Color, points []point.Id) Path {
	return Path{
		Id:     id,
		Tool:   tool,
		Color:  color,
		Points: points,
	}
}
