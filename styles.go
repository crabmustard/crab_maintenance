package main

import "github.com/charmbracelet/lipgloss"

var (
	red        = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo     = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green      = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
	crabYellow = lipgloss.AdaptiveColor{Light: "#9c9038", Dark: "#756a17"}
	crabRed    = lipgloss.AdaptiveColor{Light: "#8c4138", Dark: "#6b2820"}
	crabGreen  = lipgloss.AdaptiveColor{Light: "#1a4a0b", Dark: "#328a17"}
)

type Styles struct {
	Base,
	Menu,
	HeaderText,
	Status,
	StatusHeader,
	Highlight,
	ErrorHeaderText,
	PtacRed,
	PtacGreen,
	PtacYellow,
	Help lipgloss.Style
}

func maintStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().Foreground(indigo).Bold(true).Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(crabGreen).
		Background(crabRed).
		PaddingLeft(1).
		MarginTop(2)
	s.Menu = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(crabGreen).
		Background(crabYellow).
		PaddingLeft(1).
		MarginTop(2)
	s.StatusHeader = lg.NewStyle().
		Foreground(green).
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.
		Foreground(red)
	s.Help = lg.NewStyle().
		Foreground(lipgloss.Color("240"))
	s.PtacGreen = lg.NewStyle().
		Foreground(crabGreen)
	s.PtacRed = lg.NewStyle().
		Foreground(crabRed)
	s.PtacYellow = lg.NewStyle().
		Foreground(crabYellow)
	return &s
}
