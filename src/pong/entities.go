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
	MaxScore     = 9
	User         = 1
	ScreenHeight = 20
	ScreenWidth  = 71
	PlayAI       = `+-----------------------------Score:00|00-----------------------------+
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
`
	Multiplayer = `+-----------------------------Score:00-00-----------------------------+
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
+---------------------------------------------------------------------+`
	Menu = `+---------------------------------------------------------------------+
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                     ____   ___  _   _  ____                         |
|                    |  _ \ / _ \| \ | |/ ___|                        |
|                    | |_) | | | |  \| | |  _                         |
|                    |  __/| |_| | |\  | |_| |                        |
|                    |_|    \___/|_| \_|\____|                        |
|                         [Q] Singleplayer                            |
|                         [W] Multiplayer                             |
|                         [ESC] Quit                                  |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
+---------------------------------------------------------------------+`
	GameOver = `+---------------------------------------------------------------------+
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|         ____    _    __  __ _____    _____     _______ ____         |
|        / ___|  / \  |  \/  | ____|  / _ \ \   / / ____|  _ \        |
|       | |  _  / _ \ | |\/| |  _|   | | | \ \ / /|  _| | |_) |       |
|       | |_| |/ ___ \| |  | | |___  | |_| |\ V / | |___|  _ <        |
|        \____/_/   \_\_|  |_|_____|  \___/  \_/  |_____|_| \_\       |
|                          Score:%v-%v                                  |
|                          [ESC] Exit to the menu                     |
|                          [A] Play again                             |
|                                                                     |
|                                                                     |
|                                                                     |	
|                                                                     |
|                                                                     |
+---------------------------------------------------------------------+`
)

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
			s[i][j] = PlayAI[i*72+j]
		}
	}
	return &s
}

func changeSlope() int {
	return -1 + rand.Intn(3)
}

func (g Game) String() string {
	ans := ""
	for i := 0; i < ScreenHeight; i++ {
		for j := 0; j < ScreenWidth; j++ {
			//add score to the screen
			if g.State == PlayAI && i == 0 && g.Screen[i][j] == ':' {
				ans += fmt.Sprintf(":0%d-0%d", g.Score[0], g.Score[1])
				j += 5
			} else {
				ans += string(g.Screen[i][j])
			}
		}
		ans += "\n"
	}
	return ans
}

func stringToScreen(s string) (Screen, error) {
	var screen Screen
	for i := 0; i < ScreenHeight; i++ {
		for j := 0; j < ScreenWidth; j++ {

		}
	}
	return screen, nil
}
