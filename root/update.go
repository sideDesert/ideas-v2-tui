package root

import (
	"sideDesert/ideasv2/components"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	tabsPanel     = 0
	listPanel     = 1
	viewportPanel = 2
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle form mode first - don't let other components process keys
	cmds := make([]tea.Cmd, 0)
	if m.mode == Edit {
		m.reloadListData()
		m.mode = Read
	}

	if m.mode == Write {
		if keyMsg, ok := msg.(tea.KeyMsg); ok {
			if keyMsg.String() == "esc" {
				m.mode = Read
				return m, nil
			}
		}

		cmd, state := m.getManager().HandleUpdateForm(msg)
		if state == SaveAndExit {
			m.getManager().SaveLatestFile()
			m.mode = Read
		}
		if state == Exit {
			m.mode = Read
		}
		cmds = append(cmds, cmd)
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

	m, cmd = m.handleKeyEvent(msg)
	cmds = append(cmds, cmd)

	m.updateStyles()

	return m, tea.Batch(cmds...)
}

func (m *model) updateStyles() {
	switch m.activePanel {
	case tabsPanel:
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

	case listPanel:
		m.getManager().ListDelegate.Styles.SelectedTitle = m.theme.list.selectedTitle
		m.getManager().ListDelegate.Styles.SelectedTitle = m.theme.list.selectedTitle
		m.getManager().ListDelegate.Styles.SelectedDesc = m.theme.list.selectedDesc
		m.getManager().ListDelegate.Styles.NormalTitle = m.theme.list.normalTitle
		m.getManager().ListDelegate.Styles.NormalDesc = m.theme.list.normalDesc

		m.ProjectManager.ListDelegate.Styles.SelectedTitle = m.theme.list.selectedTitle
		m.ProjectManager.ListDelegate.Styles.SelectedDesc = m.theme.list.selectedDesc
		m.ProjectManager.ListDelegate.Styles.NormalTitle = m.theme.list.normalTitle
		m.ProjectManager.ListDelegate.Styles.NormalDesc = m.theme.list.normalDesc

		m.BookManager.ListDelegate.Styles.SelectedTitle = m.theme.list.selectedTitle
		m.BookManager.ListDelegate.Styles.SelectedDesc = m.theme.list.selectedDesc
		m.BookManager.ListDelegate.Styles.NormalTitle = m.theme.list.normalTitle
		m.BookManager.ListDelegate.Styles.NormalDesc = m.theme.list.normalDesc

	case viewportPanel:
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
	viewport, cmd := m.getManager().Viewport.Update(msg)
	cmds = append(cmds, cmd)
	m.getManager().Viewport = viewport
	if m.activePanel == listPanel {
		lm, cmd := m.getManager().List.Update(msg)
		cmds = append(cmds, cmd)
		m.getManager().List = lm
	}

	return cmds
}
