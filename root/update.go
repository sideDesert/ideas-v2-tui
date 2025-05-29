package root

import (
	"sideDesert/ideasv2/components"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
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
		case 0:
			cmd := m.HandleUpdateIdeaForm(msg)
			cmds = append(cmds, cmd)

		// Handle Book form updates
		case 2:
			cmd := m.HandleUpdateBookForm(msg)
			cmds = append(cmds, cmd)
		}

		return m, tea.Batch(cmds...)
	}

	// Only update other components when not in Write mode
	tabsModel, cmd := m.Tabs.Update(msg)
	cmds = append(cmds, cmd)

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

	m.Tabs = tabsModel.(components.TabModel)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		m.Tabs.Update(msg)
		isFilterNotInFocus := m.mode == Read &&
			m.IdeaManager.List.FilterState() != list.Filtering &&
			m.BookManager.List.FilterState() != list.Filtering

		switch {
		case key.Matches(msg, m.keys.Down):
			n := len(m.Tabs.Tabs)
			if m.activePanel == 0 && m.mode == Read {
				m.Tabs.ActiveTab = (m.Tabs.ActiveTab + 1) % n
			}

		case key.Matches(msg, m.keys.Up):
			n := len(m.Tabs.Tabs)
			if m.activePanel == 0 && m.mode == Read {
				if m.Tabs.ActiveTab == 0 {
					m.Tabs.ActiveTab = n - 1
				} else {
					m.Tabs.ActiveTab = (m.Tabs.ActiveTab - 1)
				}
			}

		case key.Matches(msg, m.keys.DeleteItem):
			if m.mode == Read && m.activePanel == titlesPanel && isFilterNotInFocus {
				switch m.Tabs.ActiveTab {
				case 0:
					i := m.IdeaManager.List.Index()
					m.IdeaManager.List.RemoveItem(i)
					m.IsTouched = true
				case 2:
					i := m.BookManager.List.Index()
					m.BookManager.List.RemoveItem(i)
					m.IsTouched = true
				}
			}

		case key.Matches(msg, m.keys.Help):
			if m.mode == Read {
				m.help.ShowAll = !m.help.ShowAll
			}

		case key.Matches(msg, m.keys.AddMode) && m.activePanel == titlesPanel:
			if isFilterNotInFocus {
				m.mode = Write
				switch m.Tabs.ActiveTab {
				case 0:
					m.IdeaManager.Form = NewIdeasForm()
					return m, m.IdeaManager.Form.Init()
				case 2:
					m.BookManager.Form = NewBooksForm()
					return m, m.BookManager.Form.Init()
				case 4:
					return m, m.BookManager.Form.Init()
				default:
					return m, nil
				}
			}

		case key.Matches(msg, m.keys.EditMode):
			if m.mode == Read && m.IdeaManager.List.FilterState() != list.Filtering {
				m.mode = Edit
				return m, nil
			}

		case key.Matches(msg, m.keys.ReadMode):
			m.mode = Read
			m.activePanel = 1
			return m, nil

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
						m.SaveFiles()
					}
				}
				return m, tea.Quit
			}
		}
	}

	return m, tea.Batch(cmds...)
}
