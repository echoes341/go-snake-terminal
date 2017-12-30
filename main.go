package main

import (
	"log"
	"time"

	"github.com/gizak/termui"
)

const (
	turnStr = "100ms"
	turn    = 100
)

func main() {
	err := termui.Init()
	if err != nil {
		log.Panicln(err)
	}
	defer termui.Close()

	b := termui.NewPar("")
	b.Height = 5
	b.Width = 5
	b.TextFgColor = termui.ColorWhite
	b.BorderLabel = ""
	b.BorderFg = termui.ColorCyan

	screen := NewScreen()
	screen.SetMenu(buildMenu())
	screen.Add(b)
	screen.Render()

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
		b.Y--
		screen.Render()
	})
	termui.Handle("/sys/kbd/<down>", func(termui.Event) {
		b.Y++
		screen.Render()
	})
	termui.Handle("/sys/kbd/<left>", func(termui.Event) {
		b.X--
		screen.Render()
	})
	termui.Handle("/sys/kbd/<right>", func(termui.Event) {
		b.X++
		screen.Render()
	})

	termui.Handle("/sys/kbd", func(termui.Event) {
		// handle all other key pressing
	})

	// handle a turn
	// Register a timer whose path is /timer/XXXms and then handle it
	// !!! Due to NewTimerCh implementations, all timers MUST have /timer/XXX path
	termui.Merge("/timer/"+turnStr, termui.NewTimerCh(turn*time.Millisecond))
	termui.Handle("/timer/"+turnStr, func(e termui.Event) {
		b.X++
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
