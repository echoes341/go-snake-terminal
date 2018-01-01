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
