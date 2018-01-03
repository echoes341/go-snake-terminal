package main

import (
	"errors"

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

func (s *snake) changeDirection(d direction) {
	opposites := map[direction]direction{
		RIGHT: LEFT,
		LEFT:  RIGHT,
		UP:    DOWN,
		DOWN:  UP,
	}

	if o := opposites[d]; o != s.direction {
		s.direction = d
	}
}

func (s *snake) hits(c Coord) bool {
	for _, b := range s.body {
		if b.X == c.X && b.Y == c.Y {
			return true
		}
	}
	return false
}

func (s *snake) head() Coord {
	return s.body[len(s.body)-1]
}

func (s *snake) die() error {
	return errors.New("Game over")
}
