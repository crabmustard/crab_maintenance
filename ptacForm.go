package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type addPtacForm struct {
	// state  state
	lg     *lipgloss.Renderer
	styles *Styles
	form   *huh.Form
	width  int
}

func checkRoomNumber(roomstring string) error {
	if len(roomstring) == 3 {
		roomint, err := strconv.Atoi(roomstring)
		if err != nil {
			return errors.New("room number, not room letters")
		}
		if (101 <= roomint && roomint <= 132) || (201 <= roomint && roomint <= 236) ||
			(301 <= roomint && roomint <= 332) || (401 <= roomint && roomint <= 427) {
			return nil
		} else {
			return errors.New("no room by that number")
		}
	}
	return errors.New("enter 3 digit room number")
}

func updatePtacForm() addPtacForm {
	m := addPtacForm{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = maintStyles(m.lg)

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("room number").
				Prompt("> ").
				Key("room").
				CharLimit(3).
				Validate(checkRoomNumber),

			huh.NewSelect[string]().
				Key("brand").
				Options(huh.NewOptions("Amana", "Trane", "Hotpoint", "Distinctions")...).
				Title("Ptac Brand"),

			huh.NewSelect[string]().
				Key("model").
				Options(huh.NewOptions("12000e", "12000h", "15000e", "15000h")...).
				Title("Ptac Model"),

			huh.NewConfirm().
				Key("done").
				Title("Add Ptac?").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("welp, finish up then")
					}
					return nil
				}).
				Affirmative("Yep").
				Negative("Wait, no"),
		),
	).
		WithWidth(45).
		WithShowHelp(false).
		WithShowErrors(false)
	m.Init()
	return m
}

func (m addPtacForm) Init() tea.Cmd {
	return m.form.Init()
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (m addPtacForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Interrupt
		case "esc", "q":
			return m, tea.Quit
		case "b":
			if m.form.State == huh.StateCompleted {
				return updatePtacForm(), nil
			}
		case "m":
			if m.form.State == huh.StateCompleted {
				return InitalMenu(), nil
			}
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Quit when the form is done.
		// cmds = append(cmds, tea.Quit)
		return m, nil
	}

	return m, tea.Batch(cmds...)
}

func (m addPtacForm) View() string {
	s := m.styles

	switch m.form.State {
	case huh.StateCompleted:
		var b strings.Builder
		fmt.Fprintf(&b, "Congratulations, you’ve added a  new ptac!!!\nRoom: %s\n", m.form.GetString("room"))
		fmt.Fprintf(&b, "Brand: %s \tModel: %s\n\nPlease tell james to clean it.\n\nb - add ptac\nm - menu",
			m.form.GetString("brand"), m.form.GetString("model"))
		return s.Status.Margin(0, 1).Padding(1, 2).Width(48).Render(b.String()) + "\n"
	default:

		// Form (left side)
		v := strings.TrimSuffix(m.form.View(), "\n\n")
		form := m.lg.NewStyle().Margin(1, 0).Render(v)

		// Status (right side)
		var status string
		{
			var (
				room         string
				brand        string
				model        string
				last_service string
			)

			room = fmt.Sprintf("Room: %s\n", m.form.GetString("room"))

			brand = fmt.Sprintf("Brand: %s\n", m.form.GetString("brand"))

			model = fmt.Sprintf("Model: %s\n", m.form.GetString("model"))

			last_service = time.Now().Format("2006-01-02")

			const statusWidth = 28
			statusMarginLeft := m.width - statusWidth - lipgloss.Width(form) - s.Status.GetMarginRight()
			status = s.Status.
				Height(lipgloss.Height(form)).
				Width(statusWidth).
				MarginLeft(statusMarginLeft).
				Render(s.StatusHeader.Render("Current Build") + "\n" +
					room +
					brand +
					model +
					last_service)
		}

		errors := m.form.Errors()
		header := m.appBoundaryView("Charm Employment Application")
		if len(errors) > 0 {
			header = m.appErrorBoundaryView(m.errorView())
		}
		body := lipgloss.JoinHorizontal(lipgloss.Left, form, status)

		footer := m.appBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))
		if len(errors) > 0 {
			footer = m.appErrorBoundaryView("")
		}

		return s.Base.Render(header + "\n" + body + "\n\n" + footer)
	}
}

func (m addPtacForm) errorView() string {
	var s string
	for _, err := range m.form.Errors() {
		s += err.Error()
	}
	return s
}

func (m addPtacForm) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(indigo),
	)
}

func (m addPtacForm) appErrorBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.ErrorHeaderText.Render(text),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(red),
	)
}
