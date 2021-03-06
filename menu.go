package main

import "github.com/gizak/termui"
import "fmt"

// Menu is a menu of the game
type Menu struct {
	*termui.Table
}

func initialMenu() *Menu {
	menu := termui.NewTable()
	menu.Rows = [][]string{
		[]string{"                                        "}, //table width...
		[]string{""},
		[]string{" ~~~ GO SNAKE! ~~~ "},
		[]string{""},
		[]string{"Welcome to Snake game. :)"},
		[]string{""},
		[]string{""},
		[]string{"Press"},
		[]string{""},
		[]string{" • <Enter> to start"},
		[]string{""},
		[]string{" • n  to reset the game"},
		[]string{""},
		[]string{" • p  to pause the game"},
		[]string{""},
		[]string{" • q  to quit"},
		[]string{""},
		[]string{""},
	}
	menu.FgColor = termui.ColorWhite
	menu.BgColor = termui.ColorDefault
	menu.TextAlign = termui.AlignCenter
	menu.Separator = false
	menu.Analysis()
	menu.SetSize()
	menu.Float = termui.AlignCenter
	menu.Border = true
	return &Menu{
		Table: menu,
	}
}

// override initial menu and the set the pause text instead of initial
func (m *Menu) setPauseMenu() {
	m.Rows = [][]string{
		[]string{"                                        "}, //table width...
		[]string{""},
		[]string{" ~~~ GO SNAKE! ~~~ "},
		[]string{""},
		[]string{"Game now is paused"},
		[]string{""},
		[]string{""},
		[]string{"Press"},
		[]string{""},
		[]string{""},
		[]string{" • n to reset the game"},
		[]string{""},
		[]string{" • p again to resume the game"},
		[]string{""},
		[]string{" • q to quit"},
		[]string{""},
		[]string{""},
	}
}

// override initial menu and the set the pause text instead of initial
func (m *Menu) setOverMenu(a int) {
	m.Rows = [][]string{
		[]string{"                                        "}, //table width...
		[]string{""},
		[]string{" ~~~ GO SNAKE! ~~~ "},
		[]string{""},
		[]string{""},
		[]string{"GAME OVER"},
		[]string{""},
		[]string{""},
		[]string{fmt.Sprintf("Your score %d", a)},
		[]string{""},
		[]string{""},
		[]string{"Press"},
		[]string{""},
		[]string{" • n to reset the game"},
		[]string{""},
		[]string{" • q to quit"},
		[]string{""},
		[]string{""},
	}
}

// Buffer is for termui.Buffer interface
func (m *Menu) Buffer() termui.Buffer {
	return m.Table.Buffer()
}
