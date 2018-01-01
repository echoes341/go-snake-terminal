package main

import (
	t "github.com/gizak/termui"
)

// Arena is the game's playground
type Arena struct {
	t.Block
	Width, Height int
	snake         *snake
}

// DefaultCell of arena table
var DefaultCell = t.Cell{
	Ch: ' ',
}

// NewArena returns an empty Arena
func NewArena(width, height int) *Arena {
	a := &Arena{
		Block:  *t.NewBlock(),
		Width:  width,
		Height: height,
	}
	a.Clear()
	a.Block.Width = width + 2
	a.Block.Height = height + 2
	return a
}

func newArenaGrid(w, h int) [][]t.Cell {
	grid := make([][]t.Cell, w)
	for k := range grid {
		grid[k] = make([]t.Cell, h)
		for j := range grid[k] {
			grid[k][j] = DefaultCell
		}
	}
	return grid
}

// Clear clears inner grid
func (a *Arena) Clear() {
	a.Grid = newArenaGrid(a.Width, a.Height)
}

// SetAll sets all the cells of a Arena the same c t.Cell
func (a *Arena) SetAll(c t.Cell) {
	for i := range a.Grid {
		for j := range a.Grid[i] {
			a.Grid[i][j] = c
		}
	}
}

// Set sets a cell of the Arena
func (a *Arena) Set(x, y int, c t.Cell) {
	a.Grid[x][y] = c
}

// SetCoord sets a cell of the Arena by coordinates
func (a *Arena) SetCoord(c Coord, cell t.Cell) {
	a.Set(c.X, c.Y, cell)
}

// Buffer implements Bufferer interface.
// So Arena can be rendered
func (a *Arena) Buffer() t.Buffer {
	// Followind similar approach as in the termui source code
	buf := a.Block.Buffer()
	for y := 0; y < a.InnerHeight(); y++ {
		for x := 0; x < a.InnerWidth(); x++ {
			buf.Set(a.InnerX()+x, a.InnerY()+y, a.Grid[x][y])
		}
	}
	return buf
}
