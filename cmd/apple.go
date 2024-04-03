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

func newApple(maxWidth, maxHeight int) *Apple {
	return &Apple{
		style:    tcell.StyleDefault.Foreground(tcell.ColorRed.TrueColor()),
		symbol:   'a',
		eatan:    false,
		position: newApplePosition(maxWidth, maxHeight),
	}
}

func newApplePosition(maxWidth, maxHeight int) Position {
	min := 1
	maxWidth--
	maxHeight--
	return Position{
		x: rand.Intn(maxWidth-min) + min,
		y: rand.Intn(maxHeight-min) + min,
	}
}

func (a *Apple) updateApplePosition(isAppleEatan bool, maxWidth, maxHeight int) {
	if isAppleEatan {
		a.position = newApplePosition(maxWidth, maxHeight)
		a.eatan = !isAppleEatan
	}
}

func (a *Apple) displayApple(s tcell.Screen, maxWidth, maxHeight int) {
	if !a.eatan {
		s.SetContent(a.position.x, a.position.y, a.symbol, nil, a.style)
	}
}
