package point

type Id string

type Point struct {
	X float64
	Y float64
}

func New(x, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}
