package pong

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type TickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
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

			case "ctrl+c":
				return g, tea.Quit
			case "esc":
				g = *NewGame()
			case "w":
				g.Bats[0].changeBatPosition(-1)
			case "s":
				g.Bats[0].changeBatPosition(1)
			case "k":
				g.Bats[1].changeBatPosition(-1)
			case "m":
				g.Bats[1].changeBatPosition(1)
			}
		case TickMsg:
			g.changeBallPosition(g.Bats)
			return g, doTick()
		}
		g.changeBallPosition(g.Bats)
	case Menu:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {

			case "ctrl+c":
				return g, tea.Quit

			case "q":
				g.State = PlayAI
			case "w":
				g.State = Multiplayer
			case "e":
				g.State = CreateRoom
			case "r":
				g.State = JoinRoom
			}
		}
	case GameOver:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {

			case "ctrl+c":
				return g, tea.Quit

			case "a":
				//TODO change this
				g = *NewGame()
				g.State = PlayAI
			case "esc":
				g = *NewGame()
				g.State = Menu
			}
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
		return fmt.Sprintf(g.State, g.Score[0])
	default:
		panic("Error, No game state, or game state hasnt been added yet")
	}
	// Send the UI for rendering
	return g.Screen.String()
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
			g.Ball.slope = 1
			g.Ball.position.X--
			g.Ball.position.Y += g.Ball.slope
		case g.Ball.position.Y == 18:
			g.Ball.slope = -1
			g.Ball.position.X--
			g.Ball.position.Y += g.Ball.slope
		case g.Ball.position.X == 2 && (g.Ball.position.Y == bats[0].position || g.Ball.position.Y == bats[0].position+1 || g.Ball.position.Y == bats[0].position+2 || g.Ball.position.Y == bats[0].position-1):
			g.Score[0]++
			g.Ball.direction = Right
			g.Ball.slope = changeSlope()
			g.Ball.position.X++
			g.Ball.position.Y += g.Ball.slope
		case g.Ball.position.X == 1:
			g.State = GameOver
		default:
			g.Ball.position.X--
			g.Ball.position.Y += g.Ball.slope
		}
	} else {
		switch {
		case g.Ball.position.Y == 1:
			g.Ball.slope = 1
			g.Ball.position.X++
			g.Ball.position.Y += g.Ball.slope
		case g.Ball.position.Y == 18:
			g.Ball.slope = -1
			g.Ball.position.X++
			g.Ball.position.Y += g.Ball.slope
		case g.Ball.position.X == 68 && (g.Ball.position.Y == bats[1].position || g.Ball.position.Y == bats[1].position+1 || g.Ball.position.Y == bats[1].position+2 || g.Ball.position.Y == bats[1].position-1):
			g.Ball.direction = Left
			g.Ball.slope = changeSlope()
			g.Ball.position.X--
			g.Ball.position.Y += g.Ball.slope
		case g.Ball.position.X == 69:
			g.State = GameOver
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
