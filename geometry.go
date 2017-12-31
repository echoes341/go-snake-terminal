package main

import t "github.com/gizak/termui"

// Coord is a simple coordinate structure
type Coord struct {
	X, Y int
}

// Point is a point on the arena
type Point struct {
	Coord
	t.Cell
}

// MoveBy moves the poin by amount on X and Y axis,
// If it's beyond the arena bounds, it starts from 0
func (p *Point) MoveBy(dX, dY int, a *Arena) {
	x, y := (p.X+dX)%a.Width, (p.Y+dY)%a.Height
	if x < 0 {
		x = a.Width + x
	}
	if y < 0 {
		y = a.Height + y
	}
	// clear current cell
	a.SetCoord(p.Coord, DefaultCell)
	p.X = x
	p.Y = y
	a.SetCoord(p.Coord, p.Cell)
}
