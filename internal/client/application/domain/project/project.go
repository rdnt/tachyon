package project

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"tachyon/pkg/uuid"
)

type Project struct {
	Id      uuid.UUID
	Name    string
	OwnerId uuid.UUID
	Paths   []Path
}

type Path struct {
	Id     uuid.UUID
	Tool   string
	Color  Color
	Points []Vector2
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

type Vector2 struct {
	X float64
	Y float64
}

func New(id uuid.UUID, ownerId uuid.UUID, name string) Project {
	return Project{
		Id:      id,
		Name:    name,
		OwnerId: ownerId,
	}
}
