package main

import (
	"github.com/gdamore/tcell/v2"
)

type (
	Direction int
	Status    int
)

const (
	LEFT = Direction(iota)
	RIGHT
	UP
	DOWN
)

const (
	ALIVE = Status(iota)
	DEAD
)

type node struct {
	body     rune
	position position
	next     *node
	prev     *node
}

type position struct {
	x, y int
}

type Snake struct {
	style            tcell.Style
	head             *node
	tail             *node
	length           int
	currentDirection Direction
	status           Status
}

func NewSnake(pos position) *Snake {
	newHead := &node{
		body:     '@',
		next:     nil,
		position: pos,
		prev:     nil,
	}
	return &Snake{
		style:            tcell.StyleDefault.Background(tcell.ColorGreenYellow).Foreground(tcell.ColorRoyalBlue),
		length:           1,
		currentDirection: RIGHT,
		head:             newHead,
		tail:             newHead,
		status:           ALIVE,
	}
}

func (s *Snake) updateSnakePosition(direction Direction) {
	switch direction {
	case LEFT:
		// x--
		s.addHead(s.head.position.x-1, s.head.position.y)
		s.removeTail()
	case RIGHT:
		// x++
		s.addHead(s.head.position.x+1, s.head.position.y)
		s.removeTail()
	case UP:
		// y--
		s.addHead(s.head.position.x, s.head.position.y-1)
		s.removeTail()
	case DOWN:
		// y++
		s.addHead(s.head.position.x, s.head.position.y+1)
		s.removeTail()
	}
}

func (s *Snake) displaySnake(scr tcell.Screen) {
	head := s.head
	for head != nil {
		scr.SetContent(head.position.x, head.position.y, head.body, nil, s.style)
		head = head.prev
	}
}

func (s *Snake) checkSnakeDeath(boardW, boardH int) {
	pos := s.head.position

	if pos.x == boardW || pos.x == 0 || pos.y == boardH || pos.y == 0 {
		s.status = DEAD
	}

	head := s.head
	for head != nil {
		if head.position.x == s.head.position.x && head.position.y == s.head.position.y && head != s.head {
			s.status = DEAD
		}
		head = head.next
	}
}

func (s *Snake) changeDirection(direction Direction) {
	if s.currentDirection == direction {
		return
	}

	if s.currentDirection == LEFT && direction == RIGHT {
		return
	}

	if s.currentDirection == RIGHT && direction == LEFT {
		return
	}

	if s.currentDirection == UP && direction == DOWN {
		return
	}

	if s.currentDirection == DOWN && direction == UP {
		return
	}

	s.currentDirection = direction
}

func (s *Snake) addHead(x, y int) {
	newHead := &node{
		body: '@',
		next: nil,
		position: position{
			x,
			y,
		},
		prev: nil,
	}

	s.head.body = '0'

	s.head.next = newHead
	newHead.prev = s.head
	s.head = newHead

	s.length++
}

func (s *Snake) removeTail() {
	list := s.tail.next

	list.prev = nil
	s.tail = nil
	s.tail = list
}
