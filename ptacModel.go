package main

import (
	"context"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/crabmustard/crab_maintenance/database"
)

type ptacModel struct {
	choices  []database.Ptac
	cursor   int
	selected map[int]struct{}
}

type foundPtacs struct {
	ptacs []database.Ptac
}

func initialPtacModel() ptacModel {
	return ptacModel{
		choices:  []database.Ptac{},
		selected: make(map[int]struct{}),
	}
}

func (pt ptacModel) Init() tea.Cmd {
	return pt.loadPtacs
}

func (pt ptacModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case foundPtacs:
		pt.choices = msg.ptacs
		return pt, nil
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
			return initialRoomModel(), nil
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
		s += fmt.Sprintf("%s [%s] %s - %s\t", cursor, checked, choice.Room, choice.LastService)
		if i%4 == 3 {
			s += "\n"
		}
	}

	s += "\nPress q to quit\n"
	return s
}

func (pt ptacModel) loadPtacs() tea.Msg {

	ptacsDb, err := cfg.db.GetAllPtac(context.Background())
	if err != nil {
		log.Fatal("error with initial ptac model")
	}
	return foundPtacs{ptacs: ptacsDb}
}
