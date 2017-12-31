package main

import t "github.com/gizak/termui"

// Screen is an abstration of the terminal's screen
type Screen struct {
	isMenuVisible bool
	menu          t.Bufferer
	buff          []t.Bufferer
}

// NewScreen returns an empty Screen
func NewScreen() *Screen {
	return &Screen{}
}

// SetMenu sets the menu of the screen(\game)
func (s *Screen) SetMenu(menu t.Bufferer) {
	s.isMenuVisible = false
	s.menu = menu
}

// Add adds elements to the screen
func (s *Screen) Add(t ...t.Bufferer) {
	s.buff = append(s.buff, t...)
}

// Remove removes an element at known position
func (s *Screen) Remove(elPos int) {
	s.buff = append(s.buff[:elPos], s.buff[elPos+1:]...)
}

// Render the screen
func (s *Screen) Render() {
	t.Clear()
	t.Render(s.buff...)
	if s.isMenuVisible {
		t.Render(s.menu)
	}
}

// ShowMenu shows the menu in the screen
func (s *Screen) ShowMenu() {
	s.isMenuVisible = true
	s.Render()
}

// RemoveMenu removes the menu from the screen
func (s *Screen) RemoveMenu() {
	s.isMenuVisible = false
	s.Render()
}

// Coord is a simple coordinate structure
type Coord struct {
	X, Y int
}

// Point is a point on the arena
type Point struct {
	Coord
	t.Cell
}

// MoveTo moves the Point to a new location
// Unsafe method. Use MoveBy instead.
func (p *Point) moveTo(x, y int, a *Arena) {
	a.SetCoord(p.Coord, DefaultCell)
	p.X = x
	p.Y = y
	a.SetCoord(p.Coord, p.Cell)
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
	p.moveTo(x, y, a)
}
