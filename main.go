package main

import (
	"log"

	"github.com/gizak/termui"
)

const (
	// turn duration in milliseconds
	turn = 300
)

// Snake's direction

func main() {
	err := termui.Init()
	if err != nil {
		log.Panicln(err)
	}
	defer termui.Close()

	g := NewGame()
	// At start the game is paused and shows the menu
	g.Start()

	termui.Loop() // block until StopLoop is called
}
