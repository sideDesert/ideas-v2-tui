package root

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	panelStyle = lipgloss.NewStyle()
)

func (m model) View() string {
	var style = m.theme.titleStyle
	leftPanelWidth := 45
	rightPanelWidth := m.width - leftPanelWidth - 8

	title := style.Render("DEBUGGER")
	var status string

	helpView := m.help.View(m.keys)
	mode := get_mode(m.mode)

	helpMetaview := m.theme.inputStyle.Render(fmt.Sprintf("\nDEBUGGER: lpw %d w %d  h %d mode %s tab %d panel %d\n", leftPanelWidth, m.width, m.height, mode, m.Tabs.ActiveTab, m.activePanel))

	s := "What should we buy at the market?\n\n"

	contentStyle := lipgloss.NewStyle().
		Width(m.width)

	switch m.mode {
	case Write:
		if m.mode == Write {
			// TODO Make this actually centered
			centeredPanel := lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				Width(60).
				Padding(1, 2).
				MarginLeft(35).
				MarginTop(8)

			var form string
			if m.Tabs.ActiveTab == 0 {
				form = m.IdeaManager.Form.View()
			}
			if m.Tabs.ActiveTab == 2 {
				form = m.BookManager.Form.View()
			}
			return centeredPanel.Render(form)
		}

	case Read, Edit:
		IdeaPanel0, IdeaPanel1 := m.IdeaPanels(leftPanelWidth, rightPanelWidth)
		IdeaTabContent := lipgloss.JoinHorizontal(
			lipgloss.Left,
			IdeaPanel0,
			IdeaPanel1,
		)

		BookPanel0, BookPanel1 := m.BookPanels(leftPanelWidth, rightPanelWidth)
		BooksTabContent := lipgloss.JoinHorizontal(
			lipgloss.Left,
			BookPanel0,
			BookPanel1,
		)

		ProjectsTabContent := strings.Builder{}
		DebuggerTabContent := contentStyle.Render(title + "\n\n" + s + status + "\n\n" + helpMetaview + helpView)

		m.Tabs.TabContent = []string{
			IdeaTabContent,
			ProjectsTabContent.String(),
			BooksTabContent,
			DebuggerTabContent,
		}
		return m.Tabs.View()
	}

	return "Invalid mode"
}

func (m *model) IdeaPanels(leftPanelWidth int, rightPanelWidth int) (string, string) {
	panelContent := panelStyle.Height(m.height - 9).Render(m.IdeaManager.List.View())

	Panel0 := panelStyle.Width(leftPanelWidth).Height(m.height - 7).Render(panelContent)

	textArea := &m.IdeaManager.TextArea
	// Set dimensions first
	textArea.SetWidth(rightPanelWidth - 4)
	textArea.SetHeight(m.height - 12)

	Panel1 := panelStyle.Width(rightPanelWidth).Height(m.height - 7).Render(textArea.View())

	return Panel0, Panel1
}

func (m *model) BookPanels(leftPanelWidth int, rightPanelWidth int) (string, string) {
	panelContent := panelStyle.Height(m.height - 9).Render(m.BookManager.List.View())

	Panel0 := panelStyle.Width(leftPanelWidth).Height(m.height - 7).Render(panelContent)

	textArea := &m.BookManager.TextArea
	// Set dimensions first
	textArea.SetWidth(rightPanelWidth - 4)
	textArea.SetHeight(m.height - 12)

	Panel1 := panelStyle.Width(rightPanelWidth).Height(m.height - 7).Render(textArea.View())

	return Panel0, Panel1
}
