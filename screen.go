package main

import "github.com/gizak/termui"

// Screen is an abstration of the terminal's screen
type Screen struct {
	isMenuVisible bool
	menu          termui.Bufferer
	buff          []termui.Bufferer
}

// NewScreen returns an empty Screen
func NewScreen() *Screen {
	return &Screen{}
}

// SetMenu sets the menu of the screen(\game)
func (s *Screen) SetMenu(menu termui.Bufferer) {
	s.isMenuVisible = true
	s.menu = menu
}

// Add adds elements to the screen
func (s *Screen) Add(t ...termui.Bufferer) {
	s.buff = append(s.buff, t...)
}

// Remove removes an element at known position
func (s *Screen) Remove(elPos int) {
	s.buff = append(s.buff[:elPos], s.buff[elPos+1:]...)
}

// Render the screen
func (s *Screen) Render() {
	termui.Clear()
	termui.Render(s.buff...)
	if s.isMenuVisible {
		termui.Render(s.menu)
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
