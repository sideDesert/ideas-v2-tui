package components

import (
	"sideDesert/ideasv2/colors"
	"sideDesert/ideasv2/keymap"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TabModel struct {
	Tabs       []string
	TabContent []string
	ActiveTab  int
	keyMap     keymap.KeyMap
	width      int
	height     int
}

func NewTabModel(tabs []string, tabContent []string, activeTab int) TabModel {
	return TabModel{
		Tabs:       tabs,
		TabContent: tabContent,
		ActiveTab:  activeTab,
		keyMap:     keymap.Keys,
	}
}

func (m TabModel) Init() tea.Cmd {
	return nil
}

func (m TabModel) SetWidth(w int) tea.Model {
	m.width = w
	return m
}

func (m TabModel) SetHeight(h int) tea.Model {
	m.height = h
	return m
}

func (m TabModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	numtabs := len(m.Tabs)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keyMap.NextTab):
			m.ActiveTab = (m.ActiveTab + 1) % numtabs
			return m, nil

		case key.Matches(msg, m.keyMap.PrevTab):
			m.ActiveTab = (m.ActiveTab - 1 + numtabs) % numtabs // wrap backwards safely
			return m, nil
		}
	}

	return m, nil
}

func tabBorderWithBottom(_border string) lipgloss.Border {
	split := strings.Split(_border, "")

	left := split[0]
	middle := split[1]
	right := split[2]
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	inactiveTabBorder = tabBorderWithBottom("┴─┴")
	activeTabBorder   = tabBorderWithBottom("┘ └")
	docStyle          = lipgloss.NewStyle().Padding(0, 0, 0, 0)
	highlightColor    = lipgloss.AdaptiveColor{Light: colors.Purple, Dark: colors.DarkPurple}

	// Gives padding to inner tab values
	inactiveTabStyle = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle   = inactiveTabStyle.Border(activeTabBorder, true)

	// This is inner window padding  good sane defaults
	windowStyle = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(0, 1).Border(lipgloss.NormalBorder()).UnsetBorderTop()
)

func (m TabModel) View() string {

	width := m.width - 2
	height := m.height - 6
	doc := strings.Builder{}

	var renderedTabs []string

	// Handle Render tabs
	for i, t := range m.Tabs {
		var style lipgloss.Style

		isFirst, isLast, isActive := i == 0, i == len(m.Tabs)-1, i == m.ActiveTab

		if isActive {
			style = activeTabStyle
		} else {
			style = inactiveTabStyle
		}

		border, _, _, _, _ := style.GetBorder()
		handleBorderEdgeCase(&border, isFirst, isActive, isLast)
		style = style.Border(border)

		renderedTabs = append(renderedTabs, style.Render(t))
	}

	// This basically places the tabs next to each other
	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	remainingRow := width - lipgloss.Width(row) + 1
	l := lipgloss.NewStyle().Foreground(lipgloss.Color(colors.Purple))

	if remainingRow > 0 {
		row += strings.Repeat(l.Render("─"), remainingRow)
	}
	row += l.Render("┐")

	doc.WriteString(row)
	doc.WriteString("\n")

	docString := windowStyle.Width(width).Height(height).Render(m.TabContent[m.ActiveTab])

	doc.WriteString(docString)
	doc.WriteString("\n")

	return docStyle.Render(doc.String())
}

func handleBorderEdgeCase(border *lipgloss.Border, isFirst bool, isActive bool, isLast bool) {
	if isFirst && isActive {
		border.BottomLeft = "│"
	} else if isFirst && !isActive {
		border.BottomLeft = "├"
	} else if isLast && isActive {
		border.BottomRight = "└"
	} else if isLast && !isActive {
		border.BottomRight = "┴"
	}
}
