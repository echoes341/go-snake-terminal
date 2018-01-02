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

func (s *snake) changeDirection(d direction) {
	opposites := map[direction]direction{
		RIGHT: LEFT,
		LEFT:  RIGHT,
		UP:    DOWN,
		DOWN:  UP,
	}

	if o := opposites[d]; o != 0 && o != s.direction {
		s.direction = d
	}
}

func (s *snake) head() (Coord, int) {
	return s.body[len(s.body)-1], len(s.body) - 1
}

func (s *snake) die() error {
	return errors.New("Game over")
}

func (s *snake) move() error {
	h, _ := s.head()
	x := h.X
	y := h.Y

	switch s.direction {
	case RIGHT:
		x++
	case LEFT:
		x--
	case UP:
		y--
	case DOWN:
		y++
	}

	s.last = s.body[0]
	s.body = append(s.body[1:], Coord{X: x, Y: y})
	return nil
}
