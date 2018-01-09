package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gizak/termui"
)

// snake directions

const (
	initDuration = 400
	minDuration  = 130
)

var debug bool

// Game is general game struct
type Game struct {
	score        int
	scoreDisplay *termui.Par
	arena        *Arena
	isStarted    bool
	IsOver       bool
	IsPaused     bool
	menu         *Menu
	PointsChan   chan int
}

// Coord is a simple coordinate structure
type Coord struct {
	X, Y int
}

// NewGame returns a new game
func NewGame(dbg bool) *Game {
	debug = dbg
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
		if g.IsOver {
			g.menu.setOverMenu(g.score)
		}
		termui.Render(g.menu)
	} else {
		// update score text
		g.scoreDisplay.Text = fmt.Sprintf("~~~ GO SNAKE! ~~~\n\nScore: %d\n", g.score)

		termui.Render(g.arena, g.scoreDisplay)
	}
}

// Start is the initial stage of the game
func (g *Game) Start() {
	// setting random seed
	rand.Seed(time.Now().UTC().UnixNano())
	g.PointsChan = make(chan int)
	g.arena.PointsChan = g.PointsChan

	go g.pointManager()

	// Renders menu game
	// and set basic handlers
	g.Render()
	g.startTimeCounter() // don't put this inside a handle

	termui.Handle("/sys/kbd/q", func(termui.Event) {
		// press q to quit
		termui.StopLoop()
	})

	// game start
	termui.Handle("/sys/kbd/<enter>", func(termui.Event) {
		if !g.isStarted { //enter works only when game is not started
			g.isStarted = true
			g.IsPaused = false
			// change menu, clear the screen
			// and start a turn

			g.menu.setPauseMenu()
			termui.Clear()
			g.handleTurn()
		}
	})

	// game pause
	termui.Handle("/sys/kbd/p", func(termui.Event) {
		if !g.IsOver {
			g.IsPaused = !g.IsPaused

			if !g.IsPaused { // clear screen before going back to game
				termui.Clear()
			}
			g.Render()
		}
	})

	// new game
	termui.Handle("/sys/kbd/n", func(termui.Event) {
		g.arena = initialArena()
		g.arena.PointsChan = g.PointsChan
		g.menu = initialMenu()
		g.score = 0
		g.IsPaused = true
		g.isStarted = false
		g.IsOver = false
		termui.Clear()
		g.Render()
	})

	//handling movements
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

	// handling turn: game's core
	termui.Handle("/timer/turn", func(termui.Event) {
		g.handleTurn()
	})
}

func (g *Game) handleTurn() {
	if g.isStarted { //turn works only when game is started
		if g.IsPaused == false { //game is not paused
			// GAME TURN IS HERE
			err := g.arena.moveSnake()
			if err != nil {
				g.IsOver = true
				g.IsPaused = true
			} else {
				termui.Clear() // clear only if game is not over
			}
		}
		g.Render()
	}
	if debug {
		g.scoreDisplay.Text = fmt.Sprintf("~~~ GO SNAKE! ~~~\n\nScore: %d\n", g.score)
		g.scoreDisplay.Text += fmt.Sprintf("Status:\nisPaused: %v\nIsOver: %v\nIsStarted: %v", g.IsPaused, g.IsOver, g.isStarted)
		termui.Render(g.scoreDisplay)
	}
}

func (g *Game) pointManager() {
	var points int
	for {
		points = <-g.PointsChan
		g.score += points
	}
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
		Coord{X: 1, Y: 4},
		Coord{X: 1, Y: 3},
		Coord{X: 1, Y: 2},
		Coord{X: 1, Y: 1},
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
