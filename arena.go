package main

import (
	t "github.com/gizak/termui"
)

// Arena is the game's playground
type Arena struct {
	t.Block
	Width, Height int
	snake         *snake
	init          bool
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
		init:   true,
		snake:  s,
	}
	a.Block.Width = width + 2
	a.Block.Height = height + 2
	return a
}

// Buffer implements Bufferer interface.
// So Arena can be rendered
func (a *Arena) Buffer() t.Buffer {
	// Followind similar approach as in the termui source code
	buf := a.Block.Buffer()

	if a.init {
		for y := 0; y < a.InnerHeight(); y++ {
			for x := 0; x < a.InnerWidth(); x++ {
				buf.Set(a.InnerX()+x, a.InnerY()+y, DefaultCell)
			}
		}
		a.init = false
	} else {
		//print snake
		for b := 0; b < a.snake.length; b++ {
			buf.Set(a.InnerX()+a.snake.body[b].X, a.InnerY()+a.snake.body[b].Y, a.snake.cell)
		}

		//set where the tail was to a defaul cell
		buf.Set(a.InnerX()+a.snake.last.X, a.InnerY()+a.snake.last.Y, DefaultCell)
	}

	return buf
}
