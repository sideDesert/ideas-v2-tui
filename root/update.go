package root

import (
	"sideDesert/ideasv2/components"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	sectionPanel = 0
	titlesPanel  = 1
	descPanel    = 2
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle form mode first - don't let other components process keys
	cmds := make([]tea.Cmd, 0)
	if m.mode == Write {
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			if keyMsg.String() == "esc" {
				m.mode = Read
				return m, nil
			}
		}
		switch m.Tabs.ActiveTab {
		// Handle IDEA form updates
		case ideasTab:
			cmd, state := m.IdeaManager.HandleUpdateForm(msg)
			if state == SaveAndExit {
				m.IdeaManager.SaveLatestFile()
				m.mode = Read
			}
			if state == Exit {
				m.mode = Read
			}

			cmds = append(cmds, cmd)
		// Handle Project form updates
		case projectsTab:
			cmd, state := m.ProjectManager.HandleUpdateForm(msg)
			if state == SaveAndExit {
				m.ProjectManager.SaveLatestFile()
				m.mode = Read
			}
			if state == Exit {
				m.mode = Read
			}
			cmds = append(cmds, cmd)

		// Handle Book form updates
		case booksTab:
			cmd, state := m.BookManager.HandleUpdateForm(msg)
			if state == SaveAndExit {
				m.BookManager.SaveLatestFile()
				m.mode = Read
			}
			if state == Exit {
				m.mode = Read
			}
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)
	}

	// Only update other components when not in Write mode
	tabsModel, cmd := m.Tabs.Update(msg)
	cmds = append(cmds, cmd)

	if tabModel, ok := tabsModel.(components.TabModel); ok {
		m.Tabs = tabModel
	}

	tcmds := m.updateTabState(&msg)
	cmds = append(cmds, tcmds...)

	m, hcmds := m.handleKeyEvent(msg)
	cmds = append(cmds, hcmds...)

	m.updateStyles()

	return m, tea.Batch(cmds...)
}

func (m *model) updateStyles() {
	switch m.activePanel {
	case 0:
		m.IdeaManager.ListDelegate.Styles.SelectedTitle = m.theme.list.selectedTitle
		m.IdeaManager.ListDelegate.Styles.SelectedDesc = m.theme.list.selectedDesc
		m.IdeaManager.ListDelegate.Styles.NormalTitle = m.theme.list.blurTitle
		m.IdeaManager.ListDelegate.Styles.NormalDesc = m.theme.list.blurDesc

		m.ProjectManager.ListDelegate.Styles.SelectedTitle = m.theme.list.selectedTitle
		m.ProjectManager.ListDelegate.Styles.SelectedDesc = m.theme.list.selectedDesc
		m.ProjectManager.ListDelegate.Styles.NormalTitle = m.theme.list.blurTitle
		m.ProjectManager.ListDelegate.Styles.NormalDesc = m.theme.list.blurDesc

		m.BookManager.ListDelegate.Styles.SelectedTitle = m.theme.list.selectedTitle
		m.BookManager.ListDelegate.Styles.SelectedDesc = m.theme.list.selectedDesc
		m.BookManager.ListDelegate.Styles.NormalTitle = m.theme.list.blurTitle
		m.BookManager.ListDelegate.Styles.NormalDesc = m.theme.list.blurDesc

	case 1:
		m.IdeaManager.ListDelegate.Styles.SelectedTitle = m.theme.list.selectedTitle
		m.IdeaManager.ListDelegate.Styles.SelectedDesc = m.theme.list.selectedDesc
		m.IdeaManager.ListDelegate.Styles.NormalTitle = m.theme.list.normalTitle
		m.IdeaManager.ListDelegate.Styles.NormalDesc = m.theme.list.normalDesc

		m.ProjectManager.ListDelegate.Styles.SelectedTitle = m.theme.list.selectedTitle
		m.ProjectManager.ListDelegate.Styles.SelectedDesc = m.theme.list.selectedDesc
		m.ProjectManager.ListDelegate.Styles.NormalTitle = m.theme.list.normalTitle
		m.ProjectManager.ListDelegate.Styles.NormalDesc = m.theme.list.normalDesc

		m.BookManager.ListDelegate.Styles.SelectedTitle = m.theme.list.selectedTitle
		m.BookManager.ListDelegate.Styles.SelectedDesc = m.theme.list.selectedDesc
		m.BookManager.ListDelegate.Styles.NormalTitle = m.theme.list.normalTitle
		m.BookManager.ListDelegate.Styles.NormalDesc = m.theme.list.normalDesc
	case 2:
		m.IdeaManager.ListDelegate.Styles.SelectedTitle = m.theme.list.selectedTitle
		m.IdeaManager.ListDelegate.Styles.SelectedDesc = m.theme.list.selectedDesc
		m.IdeaManager.ListDelegate.Styles.NormalTitle = m.theme.list.blurTitle
		m.IdeaManager.ListDelegate.Styles.NormalDesc = m.theme.list.blurDesc

		m.ProjectManager.ListDelegate.Styles.SelectedTitle = m.theme.list.selectedTitle
		m.ProjectManager.ListDelegate.Styles.SelectedDesc = m.theme.list.selectedDesc
		m.ProjectManager.ListDelegate.Styles.NormalTitle = m.theme.list.blurTitle
		m.ProjectManager.ListDelegate.Styles.NormalDesc = m.theme.list.blurDesc

		m.BookManager.ListDelegate.Styles.SelectedTitle = m.theme.list.selectedTitle
		m.BookManager.ListDelegate.Styles.SelectedDesc = m.theme.list.selectedDesc
		m.BookManager.ListDelegate.Styles.NormalTitle = m.theme.list.blurTitle
		m.BookManager.ListDelegate.Styles.NormalDesc = m.theme.list.blurDesc
	}
}

func (m *model) updateTabState(msg *tea.Msg) []tea.Cmd {
	cmds := make([]tea.Cmd, 0)
	switch m.Tabs.ActiveTab {
	case 0:
		ideaViewportModel, cmd := m.IdeaManager.Viewport.Update(msg)
		cmds = append(cmds, cmd)
		m.IdeaManager.Viewport = ideaViewportModel
		if m.activePanel == titlesPanel {
			listModel, cmd := m.IdeaManager.List.Update(msg)
			cmds = append(cmds, cmd)
			m.IdeaManager.List = listModel
		}
	case 2:
		bookViewportModel, cmd := m.BookManager.Viewport.Update(msg)
		cmds = append(cmds, cmd)
		m.BookManager.Viewport = bookViewportModel
		if m.activePanel == titlesPanel {
			listModel, cmd := m.BookManager.List.Update(msg)
			cmds = append(cmds, cmd)
			m.BookManager.List = listModel
		}
	case 1:
		projectViewportModel, cmd := m.ProjectManager.Viewport.Update(msg)
		cmds = append(cmds, cmd)
		m.ProjectManager.Viewport = projectViewportModel
		if m.activePanel == titlesPanel {
			listModel, cmd := m.ProjectManager.List.Update(msg)
			cmds = append(cmds, cmd)
			m.ProjectManager.List = listModel
		}
	}

	return cmds
}
