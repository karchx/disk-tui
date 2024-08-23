package input

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	textInput textinput.Model
}

func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Enter password"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	ti.EchoMode = textinput.EchoPassword
	ti.EchoCharacter = '*'

	return Model{
		textInput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, nil
		}

		m.textInput, cmd = m.textInput.Update(msg)
	}

	return m, cmd
}

func (m Model) View() string {
	return m.textInput.View() + "\n"
}
