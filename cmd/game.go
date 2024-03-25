package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

type State int

const (
	START = State(iota)
	RUNNING
	GAMEOVER
)

type Game struct {
	done     chan struct{}
	snake    *Snake
	board    *Board
	screen   tcell.Screen
	tickRate int
	state    State
}

func NewGame(boardHeight, boardWidth, boardOffsety, boardOffsetx int) *Game {
	screen, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e := screen.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	defStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)
	screen.SetStyle(defStyle)

	snakeStartingPos := position{
		x: boardWidth / 2,
		y: (boardHeight / 4) + boardHeight/2,
	}
	snake := NewSnake(snakeStartingPos)

	snake.addHead(snakeStartingPos.x-1, snakeStartingPos.y)
	snake.addHead(snakeStartingPos.x-2, snakeStartingPos.y)
	snake.addHead(snakeStartingPos.x-3, snakeStartingPos.y)

	return &Game{
		done:     make(chan struct{}),
		screen:   screen,
		snake:    snake,
		board:    NewBoard(boardHeight, boardWidth, boardOffsety, boardOffsetx),
		tickRate: 10,
		state:    START,
	}
}

func (g *Game) Output() {
	w, h := g.screen.Size()

	if w < g.board.w || h < g.board.h {
		g.board.DisplayScreenToSmall(g.screen, w, h)
	} else {
		switch g.state {
		case START:
			g.screen.Sync()
			g.screen.Clear()
			g.board.StartScreen(g.screen)
			g.screen.Show()
		case RUNNING:
			g.screen.Sync()
			g.screen.Clear()
			g.board.RunningScreen(g.screen, g.snake)
			g.screen.Show()
		case GAMEOVER:
			g.screen.Sync()
			g.screen.Clear()
			g.board.GameOverScreen(g.screen)
			g.screen.Show()
		}
	}
}

func (g *Game) Restart() {
	snakeStartingPos := position{
		x: g.board.w / 2,
		y: (g.board.h / 4) + g.board.h/2,
	}
	snake := NewSnake(snakeStartingPos)

	snake.addHead(snakeStartingPos.x-1, snakeStartingPos.y)
	snake.addHead(snakeStartingPos.x-2, snakeStartingPos.y)
	snake.addHead(snakeStartingPos.x-3, snakeStartingPos.y)

	g.state = RUNNING
	g.snake = snake
}

func (g *Game) gameLooop() {
	tickInterval := time.Second / time.Duration(g.tickRate)

	ticker := time.NewTicker(tickInterval)

	for {
		select {
		case <-g.done:
			ticker.Stop()
			return
		case <-ticker.C:
			if g.snake.status == DEAD {
				g.state = GAMEOVER
			}
			g.Output()
		}
	}
}
