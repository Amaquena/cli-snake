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
	apple    *Apple
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

	snakeStartingPos := Position{
		x: boardWidth/2 - 4,
		y: (boardHeight / 4) + (boardHeight / 2),
	}
	snake := NewSnake(snakeStartingPos)

	snake.addHead(snakeStartingPos.x+1, snakeStartingPos.y)
	snake.addHead(snakeStartingPos.x+2, snakeStartingPos.y)
	snake.addHead(snakeStartingPos.x+3, snakeStartingPos.y)

	return &Game{
		done:     make(chan struct{}),
		screen:   screen,
		snake:    snake,
		board:    NewBoard(boardHeight, boardWidth, boardOffsety, boardOffsetx),
		apple:    newApple(boardWidth, boardHeight),
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
			g.StartScreen()
			g.screen.Show()
		case RUNNING:
			g.screen.Sync()
			g.screen.Clear()
			g.RunningScreen()
			g.screen.Show()
		case GAMEOVER:
			g.screen.Sync()
			g.screen.Clear()
			g.GameOverScreen()
			g.screen.Show()
		}
	}
}

func (g *Game) StartScreen() {
	g.board.displayBoard(g.screen)

	style := tcell.StyleDefault.Foreground(tcell.ColorWhite.TrueColor()).Background(tcell.ColorReset)
	title := "CLI Snake"
	startButtonText := "Enter <CR> - Start Game"
	Exit := "Escape <ESC> - Exit Game"

	emitStr(g.screen, (g.board.w-len(title))/2, g.board.h/2, style, title)
	emitStr(g.screen, (g.board.w-len(startButtonText))/2, g.board.h/2+1, tcell.StyleDefault, startButtonText)
	emitStr(g.screen, (g.board.w-len(Exit))/2, g.board.h/2+2, tcell.StyleDefault, Exit)
}

func (g *Game) RunningScreen() {
	g.board.displayBoard(g.screen)
	g.apple.displayApple(g.screen, g.board.w, g.board.h)
	isAppleEatan := g.snake.checkSnakeAteApple(g.apple.position)
	g.apple.updateApplePosition(isAppleEatan, g.board.w, g.board.h)
	g.snake.checkSnakeDeath(g.board.w, g.board.h)
	if g.snake.status == ALIVE {
		g.snake.updateSnakePositionAndGrow(g.snake.currentDirection, isAppleEatan)
	}
	g.snake.displaySnake(g.screen)
}

func (g *Game) GameOverScreen() {
	g.board.displayBoard(g.screen)
	style := tcell.StyleDefault.Foreground(tcell.ColorOrangeRed.TrueColor()).Background(tcell.ColorReset)
	GameOverTitle := "GAMEOVER"
	RestartText := "Enter <CR> - Restart Game"
	Exit := "Escape <ESC> - Exit Game"

	emitStr(g.screen, (g.board.w-len(GameOverTitle))/2, g.board.h/2, style, GameOverTitle)
	emitStr(g.screen, (g.board.w-len(RestartText))/2, g.board.h/2+1, tcell.StyleDefault, RestartText)
	emitStr(g.screen, (g.board.w-len(Exit))/2, g.board.h/2+2, tcell.StyleDefault, Exit)
}

func (g *Game) Restart() {
	snakeStartingPos := Position{
		x: g.board.w/2 - 4,
		y: (g.board.h / 4) + g.board.h/2,
	}
	snake := NewSnake(snakeStartingPos)

	snake.addHead(snakeStartingPos.x+1, snakeStartingPos.y)
	snake.addHead(snakeStartingPos.x+2, snakeStartingPos.y)
	snake.addHead(snakeStartingPos.x+3, snakeStartingPos.y)

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
