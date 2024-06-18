package pong

import (
	"bytes"
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/gosuri/uilive"
	"log"
	"time"
)

func RefreshScreen(batPosition int, ballPosition Position, s Screen, writer *uilive.Writer) error {
	var buffer bytes.Buffer
	s[batPosition][1] = '['
	s[batPosition+1][1] = '['
	s[ballPosition.Y][ballPosition.X] = '*'
	for i := 0; i < 20; i++ {
		for j := 0; j < 71; j++ {
			if _, err := fmt.Fprintf(&buffer, "%s", string(s[i][j])); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintf(&buffer, "\n"); err != nil {
			return err
		}
	}
	fmt.Fprintf(writer, buffer.String())

	time.Sleep(time.Millisecond * 20)
	return nil
}

func NewGame() *Game {
	return &Game{
		Bat:  Bat{10},
		Ball: NewBall(),
	}
}

func (g *Game) Run() {
	s := NewScreen()
	writer := uilive.New()
	writer.Start()
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	timer := make(chan struct{})
	msg := make(chan string)
	go analnayaProbka(timer)
	go pressKey(msg)
	for !gameOver {
		g.Ball.changeBallPosition(g.Bat)
		RefreshScreen(g.Bat.position, g.Ball.position, *s, writer)

		select {
		case s := <-msg:
			switch {
			case s == "w":
				g.Bat.changeBatPosition(-1)
			case s == "s":
				g.Bat.changeBatPosition(1)
			case s == "esc":
				gameOver = true
			}
		case <-timer:
			continue
		}

	}
	writer.Stop()
}

func (b *Ball) changeBallPosition(bat Bat) {
	if b.direction == Left {
		switch {
		case b.position.Y == 1:
			b.slope = 1
			b.position.X--
			b.position.Y += b.slope
		case b.position.Y == 18:
			b.slope = -1
			b.position.X--
			b.position.Y += b.slope
		case b.position.X == 2 && (b.position.Y == bat.position || b.position.Y == bat.position+1 || b.position.Y == bat.position+2 || b.position.Y == bat.position-1):
			b.direction = Right
			b.slope = changeSlope()
			b.position.X++
			b.position.Y += b.slope
		case b.position.X == 1:
			gameOver = true
		default:
			b.position.X--
			b.position.Y += b.slope
		}
	} else {
		switch {
		case b.position.Y == 1:
			b.slope = 1
			b.position.X++
			b.position.Y += b.slope
		case b.position.Y == 18:
			b.slope = -1
			b.position.X++
			b.position.Y += b.slope
		case b.position.X == 68 && (b.position.Y == bat.position || b.position.Y == bat.position+1 || b.position.Y == bat.position+2 || b.position.Y == bat.position-1):
			b.direction = Left
			b.slope = changeSlope()
			b.position.X--
			b.position.Y += b.slope
		case b.position.X == 69:
			b.direction = Left
			b.slope = changeSlope()
			b.position.X--
			b.position.Y += b.slope
		default:
			b.position.X++
			b.position.Y += b.slope
		}
	}
}

func (b *Bat) changeBatPosition(amount int) {
	if b.position+amount > 0 && b.position+amount < 18 {
		b.position += amount
	}
}

func analnayaProbka(timer chan struct{}) {
	for {
		time.Sleep(50 * time.Millisecond)
		timer <- struct{}{}
	}

}

func pressKey(msg chan string) {
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Println("ERROR HERE", err)
			panic(err)
		}
		switch {
		case char == 'w':
			msg <- "w"
		case char == 's':
			msg <- "s"
		case key == keyboard.KeyEsc:
			msg <- "esc"
		}
	}
}
