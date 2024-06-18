package pong

import (
	"math/rand"
)

const (
	Left = iota
	Right
	Up
	Down
	ScreenHeight = 20
	ScreenWidth  = 71
)

var gameOver = false

type Game struct {
	Bat   Bat
	Ball  Ball
	Score int
}

type Screen [20][71]byte

type Position struct {
	X int
	Y int
}
type Bat struct {
	position int
}

type Ball struct {
	position  Position
	slope     int
	direction int
}

func NewBall() Ball {
	return Ball{
		position:  Position{X: 35, Y: 5 + rand.Intn(10)},
		slope:     -1 + rand.Intn(2),
		direction: Left,
	}
}

func NewScreen() *Screen {
	var s Screen
	for i := 0; i < ScreenHeight; i++ {
		for j := 0; j < ScreenWidth; j++ {
			if (i == 0 && j == 0) || (i == 0 && j == ScreenWidth-1) || (j == 0 && i == ScreenHeight-1) || (i == ScreenHeight-1 && j == ScreenWidth-1) {
				s[i][j] = '+'
			} else if i == 0 || i == ScreenHeight-1 {
				s[i][j] = '-'
			} else if j == 0 || j == ScreenWidth-1 {
				s[i][j] = '|'
			} else {
				s[i][j] = ' '
			}
		}
	}
	return &s
}

func changeSlope() int {
	return -1 + rand.Intn(2)
}
