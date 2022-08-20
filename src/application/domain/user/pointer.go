package user

type Mode int

const (
	Hover Mode = iota
	Pen
	Eraser
)

type Pointer struct {
	Mode Mode
	X    float64
	Y    float64
}
