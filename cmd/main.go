package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"pong/src/pong"
)

func main() {
	p := tea.NewProgram(pong.NewGame())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
