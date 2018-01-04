package main

import (
	"math/rand"

	t "github.com/gizak/termui"
)

// Arena is the game's playground
type Arena struct {
	t.Block
	Width, Height int
	snake         *snake
	food          *food
	PointsChan    chan int
}

// DefaultCell of arena table
var DefaultCell = t.Cell{
	Ch: ' ',
}

// NewArena returns an empty Arena
func NewArena(s *snake, width, height int) *Arena {
	a := &Arena{
		Block:  *t.NewBlock(),
		Width:  width,
		Height: height,
		snake:  s,
	}
	a.Block.Width = width + 2
	a.Block.Height = height + 2
	a.placeFood()
	return a
}

func (a *Arena) placeFood() {
	var x, y int

	for {
		// find a random coordinate which
		// doesn't belong to the snake

		x = rand.Intn(a.Width)
		y = rand.Intn(a.Height)

		if !a.snake.hits(Coord{x, y}) { // do while go implementation!
			break
		}
	}
	a.food = newFood(x, y)
}

func (a *Arena) hasFood(c Coord) bool {
	return c.X == a.food.X && c.Y == a.food.Y
}

// Buffer implements Bufferer interface.
// So Arena can be rendered
func (a *Arena) Buffer() t.Buffer {
	// Followind similar approach as in the termui source code
	buf := a.Block.Buffer()

	//print snake
	for b := 0; b < len(a.snake.body); b++ {
		buf.Set(a.InnerX()+a.snake.body[b].X, a.InnerY()+a.snake.body[b].Y, a.snake.cell)
	}
	buf.Set(a.InnerX()+a.food.X, a.InnerY()+a.food.Y, t.Cell{Ch: a.food.symbol})

	return buf
}

func (a *Arena) moveSnake() error {
	s := a.snake
	h := s.head()
	x := h.X
	y := h.Y

	switch s.direction {
	case RIGHT:
		x++
		x = x % a.Width
	case LEFT:
		x--
		if x < 0 {
			x = a.Width - 1
		}
	case UP:
		y--
		if y < 0 {
			y = a.Height - 1
		}
	case DOWN:
		y++
		y = y % a.Height
	}

	c := Coord{X: x, Y: y}
	if s.hits(c) {
		return s.die()
	}

	if s.length > len(s.body) {
		s.body = append(s.body, c)
	} else {
		s.body = append(s.body[1:], c)
	}

	// check if it has eaten food
	if a.hasFood(c) {
		a.placeFood()
		a.snake.length++
		a.PointsChan <- a.food.points
	}

	return nil
}
