package main

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

type Apple struct {
	style    tcell.Style
	symbol   rune
	eatan    bool
	position Position
}

func newApple(maxWidth, maxHeight, boardOffsetX, boardOffsetY int) *Apple {
	return &Apple{
		style:    tcell.StyleDefault.Foreground(tcell.ColorRed.TrueColor()),
		symbol:   'a',
		eatan:    false,
		position: newApplePosition(maxWidth, maxHeight, boardOffsetX, boardOffsetY),
	}
}

func newApplePosition(maxWidth, maxHeight, boardOffsetX, boardOffsetY int) Position {
	minX := 1 + boardOffsetX 
    minY := 1 + boardOffsetY
	maxWidth--
	maxHeight--
	return Position{
		x: rand.Intn(maxWidth-minX) + minX,
		y: rand.Intn(maxHeight-minY) + minY,
	}
}

func (a *Apple) updateApplePosition(maxWidth, maxHeight, boardOffsetX, boardOffsetY int) {
	a.position = newApplePosition(maxWidth, maxHeight, boardOffsetX, boardOffsetY)
	a.eatan = false
}

func (a *Apple) displayApple(s tcell.Screen, maxWidth, maxHeight int) {
	if !a.eatan {
		s.SetContent(a.position.x, a.position.y, a.symbol, nil, a.style)
	}
}
