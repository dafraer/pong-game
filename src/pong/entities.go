package pong

import (
	"fmt"
	"math/rand"
)

const (
	Left = iota
	Right
	MaxScore     = 9
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
+---------------------------------------------------------------------+

`
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
|                            [Q] Play                                 |
|                           [ESC] Quit                                |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
|                                                                     |
+---------------------------------------------------------------------+

`
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
+---------------------------------------------------------------------+

`
)

type game struct {
	bats   [2]bat
	ball   ball
	score  [2]int
	state  string
	screen screen
}

func NewGame() game {
	return game{
		bats:   [2]bat{bat{10}, bat{10}},
		ball:   NewBall(),
		state:  Menu,
		screen: NewScreen(),
	}
}

type screen [20][71]byte

type position struct {
	X int
	Y int
}
type bat struct {
	position int
}

type ball struct {
	position  position
	slope     int
	direction int
}

func NewBall() ball {
	return ball{
		position:  position{X: 35, Y: 5 + rand.Intn(10)},
		slope:     -1 + rand.Intn(2),
		direction: Left,
	}
}

func NewScreen() screen {
	var s screen
	for i := 0; i < ScreenHeight; i++ {
		for j := 0; j < ScreenWidth; j++ {
			s[i][j] = PlayAI[i*72+j]
		}
	}
	return s
}

func changeSlope() int {
	return -1 + rand.Intn(3)
}

func (g *game) String() string {
	ans := ""
	for i := 0; i < ScreenHeight; i++ {
		for j := 0; j < ScreenWidth; j++ {
			//add score to the screen
			if g.state == PlayAI && i == 0 && g.screen[i][j] == ':' {
				ans += fmt.Sprintf(":0%d-0%d", g.score[0], g.score[1])
				j += 5
			} else {
				ans += string(g.screen[i][j])
			}
		}
		ans += "\n"
	}
	return ans
}
