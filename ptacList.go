package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight   = 20
	defaultWidth = 60
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type ptacItem struct {
	room         string
	brand        string
	model        string
	last_service string
}

func (i ptacItem) FilterValue() string { return i.last_service }

type ptacItemDelegate struct{}

func (d ptacItemDelegate) Height() int                             { return 1 }
func (d ptacItemDelegate) Spacing() int                            { return 0 }
func (d ptacItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d ptacItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(ptacItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%5s  -%10s  -%10s  -  %s", i.room, i.brand, i.model, i.last_service)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type ptacListModel struct {
	list     list.Model
	choice   string
	quitting bool
	styles   *Styles
	lg       *lipgloss.Renderer
}

func (m ptacListModel) Init() tea.Cmd {
	return nil
}

func (m ptacListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "esc":
			return InitalMenu(), nil

		case "enter":
			i, ok := m.list.SelectedItem().(ptacItem)
			if ok {
				m.choice = string(i.room)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ptacListModel) View() string {
	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s? Serviced.", m.choice))
	}
	if m.quitting {
		return quitTextStyle.Render("Have a good day.")
	}
	return "\n" + m.list.View()
}

func newPtacList() ptacListModel {
	lm := ptacListModel{}
	ptacs, _ := cfg.db.GetAllPtac(context.Background())
	lm.lg = lipgloss.DefaultRenderer()
	lm.styles = maintStyles(lm.lg)
	items := []list.Item{}
	now := time.Now()
	for _, p := range ptacs {
		lsd := ""
		then, err := time.Parse("2006-01-02", p.LastService)
		if err != nil {
			log.Fatal(err)
		}
		diff := now.Sub(then)
		if diff > time.Duration(time.Hour*24*365*2) {
			lsd += lm.styles.PtacRed.Render(p.LastService)
		} else if diff > time.Duration(time.Hour*24*365) {
			lsd += lm.styles.PtacYellow.Render(p.LastService)
		} else {
			lsd += lm.styles.PtacGreen.Render(p.LastService)
		}
		i := ptacItem{
			room:         p.Room,
			brand:        p.Brand,
			model:        p.Model,
			last_service: lsd,
		}
		items = append(items, i)
	}

	l := list.New(items, ptacItemDelegate{}, defaultWidth, listHeight)
	l.Title = "Current Room Ptacs"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	lm.list = l

	return lm
}

func newPtacCleaningList() ptacListModel {
	lm := ptacListModel{}
	ptacs, _ := cfg.db.GetPtacsToClean(context.Background(), 10)
	lm.lg = lipgloss.DefaultRenderer()
	lm.styles = maintStyles(lm.lg)
	items := []list.Item{}
	now := time.Now()
	for _, p := range ptacs {
		lsd := ""
		then, err := time.Parse("2006-01-02", p.LastService)
		if err != nil {
			log.Fatal(err)
		}
		diff := now.Sub(then)
		if diff > time.Duration(time.Hour*24*365*2) {
			lsd += lm.styles.PtacRed.Render(p.LastService)
		} else if diff > time.Duration(time.Hour*24*365) {
			lsd += lm.styles.PtacYellow.Render(p.LastService)
		} else {
			lsd += lm.styles.PtacGreen.Render(p.LastService)
		}
		i := ptacItem{
			room:         p.Room,
			brand:        p.Brand,
			model:        p.Model,
			last_service: lsd,
		}
		items = append(items, i)
	}

	l := list.New(items, ptacItemDelegate{}, defaultWidth, listHeight)
	l.Title = "Ptacs that need cleaning"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return ptacListModel{list: l}
}
