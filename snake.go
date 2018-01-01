package main

import (
	"github.com/gizak/termui"
)

type direction byte

// snake directions
const (
	UP direction = iota
	DOWN
	LEFT
	RIGHT
)

type snake struct {
	body      []Coord
	last      Coord
	direction direction
	length    int
	cell      termui.Cell
}

func newSnake(d direction, b []Coord) *snake {
	return &snake{
		length:    len(b),
		body:      b,
		direction: d,
		cell: termui.Cell{
			Ch: '#',
		},
	}
}
