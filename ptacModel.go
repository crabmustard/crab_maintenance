package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type ptacModel struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialPtacModel() ptacModel {
	return ptacModel{
		choices:  []string{"trane", "amana", "hotpoint", "distinctions"},
		selected: make(map[int]struct{}),
	}
}

func (pt ptacModel) Init() tea.Cmd {
	return nil
}

func (pt ptacModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return pt, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if pt.cursor > 0 {
				pt.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if pt.cursor < len(pt.choices)-1 {
				pt.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := pt.selected[pt.cursor]
			if ok {
				delete(pt.selected, pt.cursor)
			} else {
				pt.selected[pt.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return pt, nil
}

func (pt ptacModel) View() string {
	s := "Ptac Status List\n\n"

	for i, choice := range pt.choices {
		// what choice is highlighted
		cursor := " "
		if pt.cursor == i {
			cursor = ">"
		}
		// what choice is selected
		checked := " "
		if _, ok := pt.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit\n"
	return s
}
