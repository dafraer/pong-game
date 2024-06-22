package pong

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type TickMsg time.Time

// doTick defines refresh rate
func doTick() tea.Cmd {
	return tea.Tick(time.Millisecond*30, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

// Init initialises game model
func (g game) Init() tea.Cmd {
	return doTick()
}

// Update updates game struct
func (g game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch g.state {
	case PlayAI:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc":
				g = NewGame()
			case "w":
				g.bats[1].changeBatPosition(g.calculateAIBatDirection())
				g.bats[0].changeBatPosition(-1)
			case "s":
				g.bats[1].changeBatPosition(g.calculateAIBatDirection())
				g.bats[0].changeBatPosition(1)
			}
		case TickMsg:
			g.bats[1].changeBatPosition(g.calculateAIBatDirection())
			g.changeBallPosition()
			return g, doTick()
		}
	case Menu:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {

			case "esc":
				return g, tea.Quit

			case "q":
				g.state = PlayAI
			}
		case TickMsg:
			return g, doTick()
		}
	case GameOver:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "a":
				g = NewGame()
				g.state = PlayAI
			case "esc":
				g = NewGame()
				g.state = Menu
			}
		case TickMsg:
			return g, doTick()
		}
	case Multiplayer:
		//TODO Add multiplayer

	}
	return g, nil
}

// View outputs game screen to the terminal
func (g game) View() string {
	switch g.state {
	case PlayAI:
		g.refreshScreen()
	case Menu:
		return g.state
	case GameOver:
		return fmt.Sprintf(g.state, g.score[0], g.score[1])
	case Multiplayer:
		g.refreshScreen()
	default:
		panic("Error, No game state, or game state hasnt been added yet")
	}
	// Send the UI for rendering
	return g.String()
}

// RefreshScreen updates screen based on the game struct
func (g *game) refreshScreen() {
	g.screen[g.bats[0].position][1] = '['
	g.screen[g.bats[0].position+1][1] = '['
	g.screen[g.bats[1].position][69] = ']'
	g.screen[g.bats[1].position+1][69] = ']'
	g.screen[g.ball.position.Y][g.ball.position.X] = '*'
}

func (g *game) changeBallPosition() {
	if g.ball.direction == Left {
		switch {
		case g.ball.position.Y == 1:
			//If the ball hits upper border it changes direction
			fmt.Print("\a")
			g.ball.slope = 1
			g.ball.position.X--
			g.ball.position.Y += g.ball.slope
		case g.ball.position.Y == 18:
			//If the ball hits lower border it changes direction
			fmt.Print("\a")
			g.ball.slope = -1
			g.ball.position.X--
			g.ball.position.Y += g.ball.slope
		case g.ball.position.X == 2 && (g.ball.position.Y == g.bats[0].position || g.ball.position.Y == g.bats[0].position+1 || (g.ball.position.Y == g.bats[0].position+2 && g.ball.slope == -1) || (g.ball.position.Y == g.bats[0].position-1 && g.ball.slope == 1)):
			//If the ball hits the bat directly or on the corner, ball is reflected
			fmt.Print("\a")
			g.ball.direction = Right
			g.ball.slope = changeSlope()
			g.ball.position.X++
			g.ball.position.Y += g.ball.slope
		case g.ball.position.X == 1:
			//If the ball hits the wall score is increased by 1
			g.score[1]++
			g.ball = NewBall()
			g.bats = [2]bat{bat{10}, bat{10}}
			if g.score[1] >= MaxScore {
				g.state = GameOver
			}
		default:
			//Just move the ball
			g.ball.position.X--
			g.ball.position.Y += g.ball.slope
		}
	} else {
		switch {
		case g.ball.position.Y == 1:
			//If the ball hits upper border it changes direction
			fmt.Print("\a")
			g.ball.slope = 1
			g.ball.position.X++
			g.ball.position.Y += g.ball.slope
		case g.ball.position.Y == 18:
			//If the ball hits lower border it changes direction
			fmt.Print("\a")
			g.ball.slope = -1
			g.ball.position.X++
			g.ball.position.Y += g.ball.slope
		case g.ball.position.X == 68 && (g.ball.position.Y == g.bats[1].position || g.ball.position.Y == g.bats[1].position+1 || (g.ball.position.Y == g.bats[1].position+2 && g.ball.slope == -1) || (g.ball.position.Y == g.bats[1].position-1 && g.ball.slope == 1)):
			//If the ball hits the bat directly or on the corner, ball is reflected
			fmt.Print("\a")
			g.ball.direction = Left
			g.ball.slope = changeSlope()
			g.ball.position.X--
			g.ball.position.Y += g.ball.slope
		case g.ball.position.X == 69:
			//If the ball hits the wall score is increased by 1
			g.score[0]++
			g.ball = NewBall()
			g.bats = [2]bat{bat{10}, bat{10}}
			if g.score[0] >= MaxScore {
				g.state = GameOver
			}
		default:
			//Just move the ball
			g.ball.position.X++
			g.ball.position.Y += g.ball.slope
		}
	}
}

func (b *bat) changeBatPosition(amount int) {
	if b.position+amount > 0 && b.position+amount < 18 {
		b.position += amount
	}
}

func (g *game) calculateAIBatDirection() int {
	//If ball flies in the opposite direction, or it is far away, or the bat is at the same Y coordinate with the ball - do nothing
	if g.ball.direction == Left || g.ball.position.X < 60 || g.bats[1].position == g.ball.position.Y {
		return 0
	}
	//If the bat is too low for the ball, move it up
	if g.bats[1].position > g.ball.position.Y {
		return -1
	}
	//If the bat is too high for the ball, move down
	if g.bats[1].position < g.ball.position.Y {
		return 1
	}
	return 0
}
