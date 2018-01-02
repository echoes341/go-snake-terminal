package main

import (
	"fmt"
	"time"

	"github.com/gizak/termui"
)

// snake directions

const (
	initDuration = 400
	minDuration  = 200
)

// Game is general game struct
type Game struct {
	score        int
	scoreDisplay *termui.Par
	arena        *Arena
	IsOver       bool
	IsPaused     bool
	menu         *Menu
}

// NewGame returns a new game
func NewGame() *Game {
	return &Game{
		arena:        initialArena(),
		score:        0,
		scoreDisplay: initialScoreDisp(),
		menu:         initialMenu(),
		IsPaused:     true,
	}
}

// Clear clears the screen
func (g *Game) Clear() {
	termui.Clear()
}

// Render renders the scene
func (g *Game) Render() {
	if g.IsPaused {
		termui.Render(g.menu)
	} else {
		// update score text
		g.scoreDisplay.Text = fmt.Sprintf("~~~ GO SNAKE! ~~~\n\nScore: %d", g.score)
		termui.Render(g.arena, g.scoreDisplay)
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

	g.startTimeCounter() // don't put this inside a handle
	termui.Handle("/sys/kbd/<enter>", func(termui.Event) {
		g.begin()
	})
}

func (g *Game) begin() {
	// change menu
	g.IsPaused = false
	g.menu.setPauseMenu()

	// Set new handlers
	termui.ResetHandlers()
	termui.Handle("/sys/kbd/q", func(termui.Event) {
		// press q to quit
		termui.StopLoop()
	})
	termui.Handle("/sys/kbd/p", func(termui.Event) {
		g.IsPaused = !g.IsPaused

		if !g.IsPaused {
			termui.Clear()
		}
		g.Render()
	})

	termui.Handle("/sys/kbd/<up>", func(termui.Event) {
		g.arena.snake.changeDirection(UP)
	})
	termui.Handle("/sys/kbd/<down>", func(termui.Event) {
		g.arena.snake.changeDirection(DOWN)
	})
	termui.Handle("/sys/kbd/<left>", func(termui.Event) {
		g.arena.snake.changeDirection(LEFT)
	})
	termui.Handle("/sys/kbd/<right>", func(termui.Event) {
		g.arena.snake.changeDirection(RIGHT)
	})

	termui.Handle("/timer/turn", func(e termui.Event) {
		if g.IsPaused == false { //game is not paused
			termui.Clear()
			// GAME TURN IS HERE
			g.arena.safeMove()
		}
		g.Render()
	})

	termui.Clear()
	g.Render()
}

func initialArena() *Arena {
	arena := NewArena(initialSnake(), 22, 20)
	arena.X = 2
	arena.Y = 2
	arena.BorderBg = termui.ColorCyan
	return arena
}

func initialSnake() *snake {
	return newSnake(RIGHT, []Coord{
		Coord{X: 1, Y: 1},
		Coord{X: 1, Y: 2},
		Coord{X: 1, Y: 3},
		Coord{X: 1, Y: 4},
	})
}

func initialScoreDisp() *termui.Par {
	b := termui.NewPar("~~~ GO SNAKE! ~~~\n\nScore: 0")
	b.Height = 22
	b.Width = 19
	b.X = 26
	b.Y = 2
	b.TextFgColor = termui.ColorWhite
	b.BorderLabel = ""
	b.BorderFg = termui.ColorCyan
	return b
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
	/*turnStr := fmt.Sprintf("/timer/%dms", initDuration)
	termui.Merge(turnStr, termui.NewTimerCh(initDuration*time.Millisecond))
	/*termui.Handle(turnStr, func(e termui.Event) {
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

func (g *Game) startTimeCounter() {
	termui.Merge("/timer/turn", g.turnTimer())
}

// This is a timer that sends an event every X time
// based on actual score
func (g *Game) turnTimer() chan termui.Event {
	t := make(chan termui.Event)
	go func(a chan termui.Event) {
		n := uint64(0)
		for {
			var ms int
			ms = initDuration - g.score
			if ms < minDuration {
				ms = minDuration
			}
			du := time.Duration(ms) * time.Millisecond
			n++
			time.Sleep(du)
			e := termui.Event{}
			e.Type = "timer"
			e.Path = "/timer/turn"
			e.Time = time.Now().Unix()
			e.Data = termui.EvtTimer{
				Duration: du,
				Count:    n,
			}
			t <- e

		}
	}(t)
	return t
}
