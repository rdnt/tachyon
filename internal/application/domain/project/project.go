package project

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"github.com/rdnt/tachyon/internal/application/domain/user"
	"github.com/rdnt/tachyon/pkg/uuid"
)

type Id uuid.UUID

type Project struct {
	Id      Id
	Name    string
	OwnerId user.Id
	Pixels  []Pixel
}

type Pixel struct {
	Color  Color
	Coords Vector2
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
	X int
	Y int
}

func New(id Id, ownerId user.Id, name string) Project {
	return Project{
		Id:      id,
		Name:    name,
		OwnerId: ownerId,
	}
}
