package main

// snake directions
const (
	UP int = iota
	DOWN
	LEFT
	RIGHT
)

type direction int

type snake struct {
	body      []Coord
	direction direction
	length    int
}

func newSnake(d direction, b []coord) *snake {
	return &snake{
		length:    len(b),
		body:      b,
		direction: d,
	}
}
