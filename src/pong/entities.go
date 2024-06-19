package pong

import (
	"fmt"
	"math/rand"
)

const (
	Left = iota
	Right
	Up
	Down
	Multiplayer  = ""
	CreateRoom   = ""
	JoinRoom     = ""
	ScreenHeight = 20
	ScreenWidth  = 71
	Menu         = `
+-----------------------------------+
|                                   |
|                                   |
|                                   |
|                                   |
|                                   |
|      ____   ___  _   _  ____      |
|     |  _ \ / _ \| \ | |/ ___|     |
|     | |_) | | | |  \| | |  _      |
|     |  __/| |_| | |\  | |_| |     |
|     |_|    \___/|_| \_|\____|     |
|          [Q] Play AI              |
|          [W] Multiplayer          |
|          [E] Create room          |
|          [R] Join room            |
|          [CTRL+C] EXIT            |
+-----------------------------------+
`
	PlayAI = `
+-----------------------------Score:%v--------------------------------+
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
+---------------------------------------------------------------------+
[ESC] Exit
`
	GameOver = `


+-----------------------------------------------------------------+
|                                                                 |
|                                                                 |
|                                                                 |
|                                                                 |
|                                                                 |
|       ____    _    __  __ _____    _____     _______ ____       |
|      / ___|  / \  |  \/  | ____|  / _ \ \   / / ____|  _ \      |
|     | |  _  / _ \ | |\/| |  _|   | | | \ \ / /|  _| | |_) |     |
|     | |_| |/ ___ \| |  | | |___  | |_| |\ V / | |___|  _ <      |
|      \____/_/   \_\_|  |_|_____|  \___/  \_/  |_____|_| \_\     |
|                          Your score:%v                           |
|                          [ESC] Exit to the menu                 |
|                          [A] Play again                         |
|                                                                 |
|                                                                 |
+-----------------------------------------------------------------+
`
)

var gameOver = false

type Game struct {
	Bats   [2]Bat
	Ball   Ball
	Score  [2]int
	State  string
	Screen Screen
}

func NewGame() *Game {
	return &Game{
		Bats:   [2]Bat{Bat{10}, Bat{10}},
		Ball:   NewBall(),
		State:  Menu,
		Screen: *NewScreen(),
	}
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

func (s Screen) String() string {
	ans := ""
	for i := 0; i < ScreenHeight; i++ {
		for j := 0; j < ScreenWidth; j++ {
			ans += string(s[i][j])
		}
		ans += "\n"
	}
	return ans
}

func stringToScreen(s string) (Screen, error) {
	var screen Screen
	if len(s) != 1420 {
		return Screen{}, fmt.Errorf("string size is not 71x20")
	}
	for i := 0; i < ScreenHeight; i++ {
		for j := 0; j < ScreenWidth; j++ {
			if i == 0 {
				screen[i][j] = s[j]
			} else if j == 0 {
				screen[i][j] = s[i*(j+1)-1]
			} else {

			}
		}
	}
	return screen, nil
}
