package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type menuModel struct {
	lg       *lipgloss.Renderer
	styles   *Styles
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func InitalMenu() menuModel {
	menu := menuModel{
		choices:  []string{"Service Ptac", "Ptac List", "Room Info"},
		selected: make(map[int]struct{}),
	}
	menu.lg = lipgloss.DefaultRenderer()
	menu.styles = maintStyles(menu.lg)

	return menu
}

func (menu menuModel) Init() tea.Cmd {
	return nil
}

func (menu menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return menu, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if menu.cursor > 0 {
				menu.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if menu.cursor < len(menu.choices)-1 {
				menu.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			switch menu.cursor {
			case 0:
				return updatePtacForm(), nil
			case 1:
				return newPtacList(), nil
			case 2:
				return newPtacCleaningList(), nil
			default:
				return menu, nil
			}

		}

	}
	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return menu, nil
}

func (menu menuModel) View() string {
	s := "Crab Mustard Properties Maintenance Ptac Portal\n\n"

	for i, choice := range menu.choices {
		// what choice is highlighted
		cursor := " "
		if menu.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s - %s\n", cursor, choice)
	}

	s += "\nPress q to quit\n"
	return menu.styles.Menu.Render(s)
}
