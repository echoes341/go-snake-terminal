package main

import (
	"log"

	"github.com/gizak/termui"
)

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
