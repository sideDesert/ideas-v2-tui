package root

import (
	"sideDesert/ideasv2/components"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
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

	ideaTextAreaModel, cmd := m.IdeaManager.TextArea.Update(msg)
	cmds = append(cmds, cmd)

	if m.activePanel == 1 {
		ideaTextAreaModel.Focus()
	} else {
		ideaTextAreaModel.Blur()
		listModel, cmd := m.IdeaManager.List.Update(msg)
		cmds = append(cmds, cmd)
		m.IdeaManager.List = listModel
	}

	activeItem := m.IdeaManager.List.Items()[m.IdeaManager.List.Index()]
	if a, ok := activeItem.(Idea); ok {
		if m.activePanel == 0 {
			ideaTextAreaModel.SetValue(a.Description())
		}
	}

	m.Tabs = tabsModel.(components.TabModel)

	m.IdeaManager.TextArea = ideaTextAreaModel

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		m.Tabs.Update(msg)
		isValid := m.mode == Read &&
			m.IdeaManager.List.FilterState() != list.Filtering &&
			m.BookManager.List.FilterState() != list.Filtering
		switch {
		case key.Matches(msg, m.keys.Up):
			if m.mode == Read {
				if m.cursor > 0 {
					m.cursor--
				}
			}

		case key.Matches(msg, m.keys.DeleteItem):
			if m.mode == Read && m.activePanel == 0 {
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

		case key.Matches(msg, m.keys.AddMode) && m.activePanel == 0:

			if isValid {
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
			m.activePanel = 0
			return m, nil

		case key.Matches(msg, m.keys.PrevPanel) || key.Matches(msg, m.keys.NextPanel):
			if m.activePanel == 1 {
				m.activePanel = 0
			} else {
				m.activePanel = 1
			}

		case key.Matches(msg, m.keys.Quit) && m.activePanel == 0:
			// quit
			if isValid {
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
