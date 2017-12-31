package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gizak/termui"
)

const (
	// turn duration in milliseconds
	turn = 300
)

// Snake's direction
const (
	UP int = iota
	DOWN
	LEFT
	RIGHT
)

func main() {
	err := termui.Init()
	if err != nil {
		log.Panicln(err)
	}
	defer termui.Close()

	screen := NewScreen()
	screen.SetMenu(buildMenu())

	arena := NewArena(20, 5)
	arena.X = 2
	arena.Y = 2
	arena.BorderBg = termui.ColorCyan

	arena.Set(0, 2, termui.Cell{Ch: ':'})
	arena.Set(1, 2, termui.Cell{Ch: ')'})

	screen.Add(arena)
	p := &Point{
		Coord: Coord{X: 0, Y: 0},
		Cell: termui.Cell{
			Ch: '█',
			Fg: termui.ColorYellow,
		},
	}
	p2 := &Point{
		Coord: Coord{X: 0, Y: 1},
		Cell: termui.Cell{
			Ch: '█',
			Fg: termui.ColorYellow,
		},
	}
	arena.SetCoord(p.Coord, p.Cell)
	arena.SetCoord(p2.Coord, p2.Cell)
	screen.Render()
	direction := DOWN

	// handle key q pressing
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		// press q to quit
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/h", func(termui.Event) {
		screen.ShowMenu()
	})

	termui.Handle("/sys/kbd/<enter>", func(termui.Event) {
		screen.RemoveMenu()
	})

	termui.Handle("/sys/kbd/<up>", func(termui.Event) {
		direction = UP
	})
	termui.Handle("/sys/kbd/<down>", func(termui.Event) {
		direction = DOWN
	})
	termui.Handle("/sys/kbd/<left>", func(termui.Event) {
		direction = LEFT
	})
	termui.Handle("/sys/kbd/<right>", func(termui.Event) {
		direction = RIGHT
	})

	termui.Handle("/sys/kbd", func(termui.Event) {
		// handle all other key pressing
	})

	// handle a turn
	// Register a timer whose path is /timer/XXXms and then handle it
	// !!! Due to NewTimerCh implementations, all timers MUST have /timer/XXX path
	turnStr := fmt.Sprintf("/timer/%dms", turn)
	termui.Merge(turnStr, termui.NewTimerCh(turn*time.Millisecond))
	termui.Handle(turnStr, func(e termui.Event) {
		switch direction {
		case UP:
			p2.MoveBy(0, -1, arena)
			p.MoveBy(0, -1, arena)
		case DOWN:
			p2.MoveBy(0, 1, arena)
			p.MoveBy(0, 1, arena)
		case LEFT:
			p2.MoveBy(-1, 0, arena)
			p.MoveBy(-1, 0, arena)
		case RIGHT:
			p2.MoveBy(1, 0, arena)
			p.MoveBy(1, 0, arena)
		}
		screen.Render()
	})

	termui.Loop() // block until StopLoop is called
}

func buildMenu() *termui.Table {
	menu := termui.NewTable()
	menu.Rows = [][]string{
		[]string{""},
		[]string{""},
		[]string{"GO SNAKE!"},
		[]string{"Welcome to Snake game. :)"},
		[]string{""},
		[]string{""},
		[]string{"Press"},
		[]string{""},
		[]string{"- <Enter> to start"},
		[]string{""},
		[]string{"- n to restart the game"},
		[]string{""},
		[]string{"  - h to see this menu again and pause the game  "},
		[]string{""},
		[]string{"- <Esc>/<q> to quit"},
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
	return menu
}
