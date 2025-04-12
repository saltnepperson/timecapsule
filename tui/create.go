package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type createState int

const (
	stateMessage createState = iota
	stateUnlockTime
	statePassphrase
	stateComplete
)

type createModel struct {
	state createState
	message string
	unlockTime string
	passphrase string
	input textinput.Model
	doneMessage string
}

func NewCreateModel() createModel {
	ti := textinput.New()
	ti.Placeholder = "Enter the secret message, dumbass"
	ti.Focus()

	return createModel{
		state: stateMessage,
		input: ti,
	}
}

func (m createModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m createModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyEnter:
				switch m.state {
				case stateMessage:
					m.message = m.input.Value()
					m.state = stateUnlockTime
					m.input.SetValue("")
					m.input.Placeholder = "e.g., 2025-12-31T12:00"
				case stateUnlockTime:
				  m.unlockTime = m.input.Value()
				  m.state = statePassphrase
				  m.input.SetValue("")
				  m.input.Placeholder = "Enter passphrase"
				  m.input.EchoMode = textinput.EchoPassword
				  m.input.EchoCharacter = '•'
				case statePassphrase:
				  m.passphrase = m.input.Value()
				  m.state = stateComplete
				  m.doneMessage = m.saveCapsule()
				case stateComplete:
					return m, tea.Quit
				}
				case tea.KeyCtrlC, tea.KeyEsc:
					return m, tea.Quit
		}
		case tea.WindowSizeMsg:
			// Do something later
			fmt.Println("Window sizing ocurred")
	}

	m.input, cmd = m.input.Update(msg)
	return m, cmd
}


func (m createModel) View() string {
	switch m.state {
	case stateMessage:
		return fmt.Sprintf("Create a Time Capsule\n\nSecret Message:\n%s", m.input.View())
	case stateUnlockTime:
		return fmt.Sprintf("Set Unlock Time (RFC3339 format, UTC)\n\n%s", m.input.View())
	case statePassphrase:
		return fmt.Sprintf("Enter a Passphrase (will be used to lock the capsule)\n\n%s", m.input.View())
	case stateComplete:
		return fmt.Sprintf("%s\n\nPress Enter to return.", m.doneMessage)
	default:
		return "Unknown state"
	}
}

func (m createModel) saveCapsule() string {
	// This is stubbed for now. We’ll save encrypted file in the next step.
	filename := fmt.Sprintf("capsule_%d.json", time.Now().Unix())
	// TODO: encrypt + save to file
	return fmt.Sprintf("✅ Capsule saved as %s (simulated)", filename)
}
