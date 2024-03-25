package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {
	const boardHeight = 30
	const boardWidth = 50
	const boardOffsetx = 0
	const boardOffsety = 0

	game := NewGame(boardHeight, boardWidth, boardOffsety, boardOffsetx)

	go game.gameLooop()

	for {
		switch ev := game.screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Output()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				game.screen.Fini()
				game.done <- struct{}{}
				os.Exit(0)
			}
			if game.state == START && ev.Key() == tcell.KeyEnter {
				game.state = RUNNING
			}
			if game.state == GAMEOVER && ev.Key() == tcell.KeyEnter {
				game.Restart()
			}
			if game.state == RUNNING {
				if ev.Key() == tcell.KeyLeft || ev.Rune() == 'a' || ev.Rune() == 'A' {
					game.snake.changeDirection(LEFT)
				}
				if ev.Key() == tcell.KeyRight || ev.Rune() == 'd' || ev.Rune() == 'D' {
					game.snake.changeDirection(RIGHT)
				}
				if ev.Key() == tcell.KeyUp || ev.Rune() == 'w' || ev.Rune() == 'W' {
					game.snake.changeDirection(UP)
				}
				if ev.Key() == tcell.KeyDown || ev.Rune() == 's' || ev.Rune() == 'S' {
					game.snake.changeDirection(DOWN)
				}
			}
		}
	}
}
