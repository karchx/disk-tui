package input

import (
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/gothew/l-og"
	cmdk "github.com/karchx/disk-tui/cmd"
)

type Model struct {
	textInput textinput.Model
	context   chan string
}

func NewModel(context chan string, placeholder string) Model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	// Many confi echo model ?
	//	ti.EchoMode = textinput.EchoPassword
	//	ti.EchoCharacter = '*'

	return Model{
		textInput: ti,
		context:   context,
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
			path := m.textInput.Value()
			device := <-m.context
			cli := cmdk.NewCommand(cmdk.Commands{
				Command: "sudo",
				Args:    []string{"mount"},
				Path:    path,
			})
			_, err := cli.MountDisk(device)
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}
			return m, nil
		}

		m.textInput, cmd = m.textInput.Update(msg)
	}

	return m, cmd
}

func (m Model) View() string {
	return m.textInput.View() + "\n"
}
