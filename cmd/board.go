package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

type Board struct {
	h, w             int
	offsetx, offsety int
}

func NewBoard(h, w, offsety, offsetx int) *Board {
	return &Board{
		h,
		w,
		offsetx,
		offsety,
	}
}

func emitStr(s tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}

func (b *Board) StartScreen(s tcell.Screen) {
	b.displayBoard(s)

	style := tcell.StyleDefault.Foreground(tcell.ColorWhite.TrueColor()).Background(tcell.ColorReset)
	title := "CLI Snake"
	startButtonText := "Enter <CR> - Start Game"
	Exit := "Escape <ESC> - Exit Game"

	emitStr(s, (b.w-len(title))/2, b.h/2, style, title)
	emitStr(s, (b.w-len(startButtonText))/2, b.h/2+1, tcell.StyleDefault, startButtonText)
	emitStr(s, (b.w-len(Exit))/2, b.h/2+2, tcell.StyleDefault, Exit)
}

func (b *Board) RunningScreen(s tcell.Screen, snake *Snake) {
	b.displayBoard(s)
	snake.updateSnakePosition(snake.currentDirection)
	snake.displaySnake(s)
	snake.checkSnakeDeath(b.w, b.h)
}

func (b *Board) GameOverScreen(s tcell.Screen) {
	b.displayBoard(s)
	style := tcell.StyleDefault.Foreground(tcell.ColorOrangeRed.TrueColor()).Background(tcell.ColorReset)
	GameOverTitle := "GAMEOVER"
	RestartText := "Enter <CR> - Restart Game"
	Exit := "Escape <ESC> - Exit Game"

	emitStr(s, (b.w-len(GameOverTitle))/2, b.h/2, style, GameOverTitle)
	emitStr(s, (b.w-len(RestartText))/2, b.h/2+1, tcell.StyleDefault, RestartText)
	emitStr(s, (b.w-len(Exit))/2, b.h/2+2, tcell.StyleDefault, Exit)
}

func (b *Board) DisplayScreenToSmall(s tcell.Screen, currentW, currentH int) {
	style := tcell.StyleDefault.Foreground(tcell.ColorIndianRed.TrueColor()).Background(tcell.Color(tcell.ColorBlack.TrueColor()))
	displayText := fmt.Sprintf("Current Screen size too small, requird size: %dw x %dh", b.w, b.h)
	currentSize := fmt.Sprintf("current Size: %dw x %dh", currentW, currentH)

	emitStr(s, (currentW-len(displayText))/2, currentH/2, style, displayText)
	emitStr(s, (currentW-len(currentSize))/2, currentH/2+1, tcell.StyleDefault, currentSize)
}

func (b *Board) displayBoard(s tcell.Screen) {
	style := tcell.StyleDefault.Foreground(tcell.ColorWhite.TrueColor()).Background(tcell.ColorReset)

	// x == w, y == h
	s.SetContent(0+b.offsetx, 0+b.offsety, tcell.RuneULCorner, nil, style)
	s.SetContent(0+b.offsetx, b.h+b.offsety, tcell.RuneLLCorner, nil, style)
	s.SetContent(b.w+b.offsetx, 0+b.offsety, tcell.RuneURCorner, nil, style)
	s.SetContent(b.w+b.offsetx, b.h+b.offsety, tcell.RuneLRCorner, nil, style)

	for x := 0; x < b.w; x++ {
		for y := 0; y < b.h; y++ {
			if y == 0 && x != 0 {
				s.SetContent(x+b.offsetx, y+b.offsety, tcell.RuneHLine, nil, style)
			}
			if x == 0 && y != 0 {
				s.SetContent(x+b.offsetx, y+b.offsety, tcell.RuneVLine, nil, style)
			}
		}
	}

	for x := 0; x <= b.w; x++ {
		for y := 0; y <= b.h; y++ {
			if x != b.w && x != 0 && y == b.h {
				s.SetContent(x+b.offsetx, y+b.offsety, tcell.RuneHLine, nil, style)
			}

			if y != b.h && y != 0 && x == b.w {
				s.SetContent(x+b.offsetx, y+b.offsety, tcell.RuneVLine, nil, style)
			}
		}
	}
}
