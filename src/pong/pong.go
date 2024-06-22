package pong

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"pong/src/udpclient"
	"time"
)

type TickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(time.Millisecond*30, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (g Game) Init() tea.Cmd {
	return doTick()
}

func (g Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch g.State {
	case PlayAI:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc":
				g = *NewGame()
			case "w":
				g.Bats[1].changeBatPosition(g.calculateAIBatDirection())
				g.Bats[0].changeBatPosition(-1)
			case "s":
				g.Bats[1].changeBatPosition(g.calculateAIBatDirection())
				g.Bats[0].changeBatPosition(1)
			}
		case TickMsg:
			g.Bats[1].changeBatPosition(g.calculateAIBatDirection())
			g.changeBallPosition(g.Bats)
			return g, doTick()
		}
	case Menu:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {

			case "esc":
				return g, tea.Quit

			case "q":
				g.State = PlayAI
			case "w":
				g.State = Multiplayer
			}
		case TickMsg:
			return g, doTick()
		}
	case GameOver:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "a":
				//TODO change this
				g = *NewGame()
				g.State = PlayAI
			case "esc":
				g = *NewGame()
				g.State = Menu
			}
		case TickMsg:
			return g, doTick()
		}
	case Multiplayer:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc":
				g = *NewGame()
			case "w":
				g = udpclient.Request(udpclient.Message{"w", User})
			case "s":
				udpclient.Request(udpclient.Message{"s", User})
			}
		case TickMsg:
			g = udpclient.Get()
			return g, doTick()
		}
	default:
		panic("Error, No game state, or game state hasnt been added yet")
	}
	return g, nil
}

func (g Game) View() string {
	switch g.State {
	case PlayAI:
		g.RefreshScreen()
	case Menu:
		return g.State
	case GameOver:
		return fmt.Sprintf(g.State, g.Score[0], g.Score[1])
	case Multiplayer:
		g.RefreshScreen()
	default:
		panic("Error, No game state, or game state hasnt been added yet")
	}
	// Send the UI for rendering
	return g.String()
}

func (g *Game) RefreshScreen() error {
	g.Screen[g.Bats[0].position][1] = '['
	g.Screen[g.Bats[0].position+1][1] = '['
	g.Screen[g.Bats[1].position][69] = ']'
	g.Screen[g.Bats[1].position+1][69] = ']'
	g.Screen[g.Ball.position.Y][g.Ball.position.X] = '*'
	return nil
}

func (g *Game) changeBallPosition(bats [2]Bat) {
	if g.Ball.direction == Left {
		switch {
		case g.Ball.position.Y == 1:
			fmt.Print("\a")
			g.Ball.slope = 1
			g.Ball.position.X--
			g.Ball.position.Y += g.Ball.slope
		case g.Ball.position.Y == 18:
			fmt.Print("\a")
			g.Ball.slope = -1
			g.Ball.position.X--
			g.Ball.position.Y += g.Ball.slope
		case g.Ball.position.X == 2 && (g.Ball.position.Y == bats[0].position || g.Ball.position.Y == bats[0].position+1 || g.Ball.position.Y == bats[0].position+2 || g.Ball.position.Y == bats[0].position-1):
			fmt.Print("\a")
			g.Ball.direction = Right
			g.Ball.slope = changeSlope()
			g.Ball.position.X++
			g.Ball.position.Y += g.Ball.slope
		case g.Ball.position.X == 1:
			g.Score[1]++
			g.Ball = NewBall()
			g.Bats = [2]Bat{Bat{10}, Bat{10}}
			if g.Score[1] >= MaxScore {
				g.State = GameOver
			}
		default:
			g.Ball.position.X--
			g.Ball.position.Y += g.Ball.slope
		}
	} else {
		switch {
		case g.Ball.position.Y == 1:
			fmt.Print("\a")
			g.Ball.slope = 1
			g.Ball.position.X++
			g.Ball.position.Y += g.Ball.slope
		case g.Ball.position.Y == 18:
			fmt.Print("\a")
			g.Ball.slope = -1
			g.Ball.position.X++
			g.Ball.position.Y += g.Ball.slope
		case g.Ball.position.X == 68 && (g.Ball.position.Y == bats[1].position || g.Ball.position.Y == bats[1].position+1 || g.Ball.position.Y == bats[1].position+2 || g.Ball.position.Y == bats[1].position-1):
			fmt.Print("\a")
			g.Ball.direction = Left
			g.Ball.slope = changeSlope()
			g.Ball.position.X--
			g.Ball.position.Y += g.Ball.slope
		case g.Ball.position.X == 69:
			g.Score[0]++
			g.Ball = NewBall()
			g.Bats = [2]Bat{Bat{10}, Bat{10}}
			if g.Score[0] >= MaxScore {
				g.State = GameOver
			}
		default:
			g.Ball.position.X++
			g.Ball.position.Y += g.Ball.slope
		}
	}
}

func (b *Bat) changeBatPosition(amount int) {
	if b.position+amount > 0 && b.position+amount < 18 {
		b.position += amount
	}
}

func (g Game) calculateAIBatDirection() int {
	if g.Ball.direction == Left || g.Ball.position.X < 60 {
		return 0
	}
	if g.Bats[1].position == g.Ball.position.Y {
		return 0
	}
	if g.Bats[1].position > g.Ball.position.Y {
		return -1
	}
	if g.Bats[1].position < g.Ball.position.Y {
		return 1
	}
	return 0
}
