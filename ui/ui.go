package ui

import (
	"github.com/karchx/disk-tui/ui/input"
	"github.com/karchx/disk-tui/ui/list"

	tea "github.com/charmbracelet/bubbletea"
)

type state int

const (
	inputState state = iota
	listState
)

type Model struct {
	state state
	list  list.Model
	input input.Model
}

func NewModel(drives []string) Model {
	m := Model{
		list:  list.NewModel(drives),
		input: input.NewModel("Path mount device (/to/path)"),
		state: listState,
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

		switch m.state {
		case inputState:
			input, cmd := m.input.Update(msg)
			m.input = input
			if msg.String() == "enter" {
				m.state = listState
			}
			return m, cmd

		case listState:
			list, cmd := m.list.Update(msg)
			m.list = list
			m.input.SetDevice(m.list.DeviceMount)

			if msg.String() == "enter" {
				m.state = inputState
			}
			return m, cmd
		}
	}

	return m, cmd
}

func (m Model) View() string {
	switch m.state {
	case inputState:
		return m.input.View()
	case listState:
		return m.list.View()
	default:
		return ""
	}
}
