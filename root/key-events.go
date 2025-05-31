package root

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

const (
	ideasTab    = 0
	projectsTab = 1
	booksTab    = 2
)

func (m model) handleKeyEvent(msg tea.Msg) (model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	if m.activePanel == listPanel {
		lm, cmd := m.getManager().List.Update(msg)
		m.getManager().List = lm
		cmds = append(cmds, cmd)
	}

	if m.activePanel == viewportPanel {
		lm, cmd := m.getManager().Viewport.Update(msg)
		lm.Height = m.height - 3
		m.getManager().Viewport = lm

		cmds = append(cmds, cmd)
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		isFilterNotInFocus := m.mode == Read &&
			m.IdeaManager.List.FilterState() != list.Filtering &&
			m.BookManager.List.FilterState() != list.Filtering

		switch {
		case key.Matches(msg, m.keys.Down):
			n := len(m.Tabs.Tabs)
			if m.activePanel == tabsPanel && m.mode == Read {
				m.Tabs.ActiveTab = (m.Tabs.ActiveTab + 1) % n
			}

			if m.activePanel == listPanel && m.mode == Read {
				li := m.getManager().GetActiveListItem()
				if li != nil {
					m.getManager().SetViewportContent(li.Description())
				}
			}
			return m, nil

		case key.Matches(msg, m.keys.Up):
			n := len(m.Tabs.Tabs)
			if m.activePanel == tabsPanel && m.mode == Read {
				if m.Tabs.ActiveTab == 0 {
					m.Tabs.ActiveTab = n - 1
				} else {
					m.Tabs.ActiveTab = (m.Tabs.ActiveTab - 1)
				}
			}

			li := m.getManager().GetActiveListItem()
			if li != nil {
				m.getManager().SetViewportContent(li.Description())
			}

			// }
			return m, nil

		case key.Matches(msg, m.keys.DeleteItem):
			if m.mode == Read && m.activePanel == listPanel && isFilterNotInFocus {
				m.mode = Delete
			}

		case key.Matches(msg, m.keys.Help):
			if m.mode == Read {
				m.help.ShowAll = !m.help.ShowAll
			}

		case key.Matches(msg, m.keys.AddMode) && m.activePanel == listPanel:
			if isFilterNotInFocus {
				m.mode = Write
				m.getManager().Form = m.NewTabForm()
				cmds = append(cmds, m.getManager().Form.Init())
				return m, tea.Batch(cmds...)
			}

		case key.Matches(msg, m.keys.EditMode):
			if m.mode == Read && isFilterNotInFocus {
				m.mode = Edit
				return m, m.editDescription()
			}

		case key.Matches(msg, m.keys.ReadMode):
			m.mode = Read
			m.activePanel = 1

		case key.Matches(msg, m.keys.NextPanel):
			m.activePanel = (m.activePanel + 1) % m.n_panels

		case key.Matches(msg, m.keys.PrevPanel):
			if m.activePanel == 0 {
				m.activePanel = 2
			} else {
				m.activePanel -= 1
			}

		case key.Matches(msg, m.keys.Quit):
			if isFilterNotInFocus {
				m.quitting = true
				if m.IsTouched {
					save := false
					huh.
						NewConfirm().
						Title("Save Changes?").
						Affirmative("Yes").
						Negative("No").
						Value(&save).
						Run()

					if save {
						m.IdeaManager.SaveLatestFile()
						m.BookManager.SaveLatestFile()
						m.ProjectManager.SaveLatestFile()
					}
				}
				cmds = append(cmds, tea.Quit)
				return m, tea.Batch(cmds...)
			}
		case key.Matches(msg, m.keys.Confirm):
			if m.mode == Delete {
				index := m.getManager().List.Index()
				m.getManager().RemoveItem(index)
				m.mode = Read
			}
		case key.Matches(msg, m.keys.Cancel):
			if m.mode == Delete {
				m.mode = Read
			}
		}

	}
	return m, tea.Batch(cmds...)
}
