package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	tui "github.com/saltnepperson/timecapsule/tui"
)

func main() {
	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
