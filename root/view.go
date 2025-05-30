package root

import (
	"sideDesert/ideasv2/colors"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

var (
	tabStyle = lipgloss.NewStyle()
)

func (m model) View() string {
	marginTop := 1
	footerHeight := 1
	if m.help.ShowAll {
		footerHeight = 4
	}
	titleHeight := 1
	tabContentHeight := m.height - marginTop - footerHeight - titleHeight

	tabsPanelWidth := m.width / 6
	titlePanelWidth := (m.width - tabsPanelWidth) / 3
	formWidth := m.width / 2
	descPanelWidth := m.width - tabsPanelWidth - titlePanelWidth

	switch m.mode {
	case Write:
		if m.mode == Write {
			// TODO Make this actually centered
			centeredPanel := lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				Width(formWidth).
				Padding(1, 2).
				MarginLeft((m.width - formWidth) / 2).
				MarginTop(8)

			var form string
			if m.Tabs.ActiveTab == ideasTab {
				form = m.IdeaManager.Form.View()
			}
			if m.Tabs.ActiveTab == projectsTab {
				form = m.ProjectManager.Form.View()
			}
			if m.Tabs.ActiveTab == booksTab {
				form = m.BookManager.Form.View()
			}

			return centeredPanel.Render(form)
		}

	case Read, Edit:
		// var tabStyle = lipgloss.NewStyle().
		// 	Margin(0).
		// 	Padding(0)

		// var footerStyle = lipgloss.NewStyle()

		var titleStyle = lipgloss.
			NewStyle().
			MarginRight(2).
			PaddingLeft(1)

		var activeTitleStyle = titleStyle.
			Background(lipgloss.Color(colors.ActiveTab)).
			Foreground(lipgloss.Color(colors.White))

		var inactiveTitleStyle = titleStyle.
			Background(lipgloss.Color(colors.InactiveTab)).
			Foreground(lipgloss.Color(colors.InactiveText))

		var contentStyle = lipgloss.NewStyle().Padding(1, 0, 0, 1)

		var tabsTitleStyle = inactiveTitleStyle
		var titlesTitleStyle = inactiveTitleStyle
		var descriptionTitleStyle = inactiveTitleStyle

		var activeTabStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(colors.ActiveText))
		var inactiveTabStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color(colors.InactiveText))

		tabsList := strings.Builder{}
		arrow := "→ "
		for i, tab := range m.Tabs.Tabs {
			textStyle := inactiveTabStyle
			text := tab

			if m.Tabs.ActiveTab == i {
				textStyle = activeTabStyle
				text = arrow + text
			}

			tabsList.WriteString(textStyle.Render(text) + "\n")
		}

		switch m.activePanel {
		case 0:
			tabsTitleStyle = activeTitleStyle
		case 1:
			titlesTitleStyle = activeTitleStyle
		case 2:
			descriptionTitleStyle = activeTitleStyle
		}

		tabsTitleStyle = tabsTitleStyle.Width(tabsPanelWidth).Height(1)
		titlesTitleStyle = titlesTitleStyle.Width(titlePanelWidth).Height(1)
		descriptionTitleStyle = descriptionTitleStyle.Width(descPanelWidth).Height(1)

		var activeMger *Manager
		switch m.Tabs.ActiveTab {
		case ideasTab:
			activeMger = &m.IdeaManager
		case projectsTab:
			activeMger = &m.ProjectManager
		case booksTab:
			activeMger = &m.BookManager
		}

		listView := &activeMger.List
		viewportView := &activeMger.Viewport
		activeIndex := activeMger.List.Index()

		if len(listView.Items()) != 0 {
			activeItem := &listView.Items()[activeIndex]
			if a, ok := (*activeItem).(Idea); ok {
				out, err := glamour.Render(a.Description(), "dark")
				if err != nil {
					tea.Println("Glamour Render error, ", err)
				}
				viewportView.SetContent(out)
			}
			listView.SetDelegate(m.IdeaManager.ListDelegate)
		}

		listView.SetHeight(tabContentHeight)
		listView.SetWidth(titlePanelWidth)
		viewportView.Height = tabContentHeight
		viewportView.Width = descPanelWidth

		return lipgloss.JoinVertical(
			lipgloss.Top,
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				lipgloss.JoinVertical(
					lipgloss.Top,
					tabsTitleStyle.
						Render("Tabs"),
					contentStyle.
						Width(tabsPanelWidth).
						Height(tabContentHeight).
						Render(tabsList.String()),
				),
				lipgloss.JoinVertical(
					lipgloss.Top,
					titlesTitleStyle.
						Render("Titles"),
					listView.View(),
				),
				lipgloss.JoinVertical(
					lipgloss.Top,
					descriptionTitleStyle.
						Render("Description"),
					viewportView.View(),
				),
			),

			m.help.View(m.keys),

			// footerStyle.Render(tabStyle.
			// 	Width(m.width).
			// 	MarginTop(1).
			// 	Height(footerHeight).
			// 	Background(lipgloss.Color("60")).
			// 	PaddingLeft(1).
			// 	Render(fmt.Sprintf("STATUS  Active Panel: %d Mode: %s", m.activePanel, get_mode(m.mode)))),
		)

	}
	return "Invalid mode"

}

func (m *model) IdeaPanels(leftPanelWidth int, rightPanelWidth int) (string, string) {
	panelContent := tabStyle.Height(m.height - 9).Render(m.IdeaManager.List.View())

	Panel0 := tabStyle.Width(leftPanelWidth).Height(m.height - 7).Render(panelContent)

	m.IdeaManager.Viewport.Width = rightPanelWidth - 4
	m.IdeaManager.Viewport.Height = m.height - 12

	Panel1 := tabStyle.Width(rightPanelWidth).Height(m.height - 7).Render(m.IdeaManager.Viewport.View())
	// Panel1 := m.IdeaManager.Viewport.View()
	return Panel0, Panel1
}

func (m *model) BookPanels(leftPanelWidth int, rightPanelWidth int) (string, string) {
	panelContent := tabStyle.Height(m.height - 9).Render(m.BookManager.List.View())

	Panel0 := tabStyle.Width(leftPanelWidth).Height(m.height - 7).Render(panelContent)

	m.BookManager.Viewport.Width = rightPanelWidth - 4
	m.BookManager.Viewport.Height = m.height - 12

	Panel1 := tabStyle.Width(rightPanelWidth).Height(m.height - 7).Render(m.IdeaManager.Viewport.View())

	return Panel0, Panel1
}
