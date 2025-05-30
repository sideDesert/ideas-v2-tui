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

func (m model) handleKeyEvent(msg tea.Msg) (model, []tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	if m.Tabs.ActiveTab == ideasTab && m.activePanel == 1 {
		lm, cmd := m.IdeaManager.List.Update(msg)
		m.IdeaManager.List = lm
		cmds = append(cmds, cmd)
	}

	if m.Tabs.ActiveTab == projectsTab && m.activePanel == 1 {
		lm, cmd := m.ProjectManager.List.Update(msg)
		m.ProjectManager.List = lm
		cmds = append(cmds, cmd)
	}

	if m.Tabs.ActiveTab == booksTab && m.activePanel == 1 {
		lm, cmd := m.BookManager.List.Update(msg)
		m.BookManager.List = lm
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
			if m.activePanel == 0 && m.mode == Read {
				m.Tabs.ActiveTab = (m.Tabs.ActiveTab + 1) % n
			}
			return m, nil

		case key.Matches(msg, m.keys.Up):
			n := len(m.Tabs.Tabs)
			if m.activePanel == 0 && m.mode == Read {
				if m.Tabs.ActiveTab == 0 {
					m.Tabs.ActiveTab = n - 1
				} else {
					m.Tabs.ActiveTab = (m.Tabs.ActiveTab - 1)
				}
			}
			return m, nil

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
					cmds = append(cmds, m.IdeaManager.Form.Init())
					return m, cmds
				case 2:
					m.BookManager.Form = NewBooksForm()
					cmds = append(cmds, m.BookManager.Form.Init())
					return m, cmds
				case 4:
					cmds = append(cmds, m.BookManager.Form.Init())
					return m, cmds
				default:
					return m, nil
				}
			}

		case key.Matches(msg, m.keys.EditMode):
			if m.mode == Read && m.IdeaManager.List.FilterState() != list.Filtering {
				m.mode = Edit
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
						m.SaveFiles()
					}
				}
				cmds = append(cmds, tea.Quit)
				return m, cmds
			}
		}
	}
	return m, cmds
}
