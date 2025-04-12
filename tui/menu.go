package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choice string
}

func InitialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "1":
			return NewCreateModel(), nil
		case "2":
			m.choice = "open"
			return m, tea.Quit
		case "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.choice != "" {
		return fmt.Sprintf("Selected: %s\n", m.choice)
	}

	return `
	Time Capsule

	1. Create a new message
	2. Read a message
	
	q. Quit this stupid application
	
	**Standard messaging rates may apply.
	`
}
