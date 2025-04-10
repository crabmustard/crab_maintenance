package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type ptacService struct {
	textInput textinput.Model
	err       error
}

// type (
// 	errMsg error
// )

func NewPtacService() ptacService {
	ti := textinput.New()
	ti.Placeholder = "- - -"
	ti.Focus()
	ti.CharLimit = 3
	ti.Width = 20

	return ptacService{
		textInput: ti,
		err:       nil,
	}
}

func (ps ptacService) Init() tea.Cmd {
	return textinput.Blink
}

func (ps ptacService) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			return InitalMenu(), nil
		}
		ps.textInput, cmd = ps.textInput.Update(msg)

		// Return the updated model to the Bubble Tea runtime for processing.
		// Note that we're not returning a command.
	}
	return ps, cmd
}

func (ps ptacService) View() string {
	s := "Crab Mustard Properties Maintenance Ptac Portal\n\n"

	s += fmt.Sprintf("What Room?\n%s", ps.textInput.View())

	s += "\nPress q to quit\n"
	return s
}
