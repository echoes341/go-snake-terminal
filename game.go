package main

import (
	"github.com/gizak/termui"
)

// snake directions
const (
	UP int = iota
	DOWN
	LEFT
	RIGHT
)

// Game is general game struct
type Game struct {
	score  int
	arena  *Arena
	IsOver bool
	menu   *Menu
}

// NewGame returns a new game
func NewGame() *Game {
	return &Game{
		arena: initialArena(),
		score: initialScore(),
		menu:  initialMenu(),
	}
}

// Clear clears the screen
func (g *Game) Clear() {
	termui.Clear()
}

// Render renders the scene
func (g *Game) Render() {
	if g.menu.IsVisible {
		termui.Render(g.menu)
	} else {
		termui.Render( /*g.arena*/ )
	}
}

// Start is the initial stage of the game
func (g *Game) Start() {
	// Renders menu game
	// and set basic handlers
	g.Render()
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		// press q to quit
		termui.StopLoop()
	})

	termui.Handle("/sys/kbd/<enter>", func(termui.Event) {
		g.begin()
	})
}

func (g *Game) begin() {
	// change menu
	g.menu.IsVisible = false
	g.menu.setPauseMenu()

	// Set new handlers
	termui.ResetHandlers()
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		// press q to quit
		termui.StopLoop()
	})
	termui.Handle("/sys/kbd/p", func(termui.Event) {
		g.menu.IsVisible = !g.menu.IsVisible
		termui.Clear()
		g.Render()
	})
	// termui.Handle()
	termui.Clear()
	g.Render()
	//g.Render()

}

func initialArena() *Arena {
	arena := NewArena(20, 20)
	arena.X = 2
	arena.Y = 2
	arena.BorderBg = termui.ColorCyan
	return arena
}

func initialScore() int {
	return 0
}

func (g *Game) initHandles() {
	// handle key q pressing

	/*
		termui.Handle("/sys/kbd/h", func(termui.Event) {
			g.ShowMenu()
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
	*/
	/* termui.Handle("/sys/kbd", func(termui.Event) {
		// handle all other key pressing
	}) */

	// handle a turn
	// Register a timer whose path is /timer/XXXms and then handle it
	// !!! Due to NewTimerCh implementations, all timers MUST have /timer/XXX path
	/* turnStr := fmt.Sprintf("/timer/%dms", turn)
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
		g.Render()
	}) */

}
