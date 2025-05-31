package root

import (
	"fmt"
	"sideDesert/ideasv2/components"
	"sideDesert/ideasv2/keymap"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

// Change this
type item string

func (m *model) reloadListData() tea.Cmd {
	data := m.getCurrentListData()
	if data == nil {
		fmt.Println("Error[*model.reloadListData]: Couldn't Load data")
		return nil
	}

	items := []list.Item{}
	for _, i := range data {
		items = append(items, i)
	}

	cmd := m.getManager().List.SetItems(items)
	return cmd
}

func (m *model) getCurrentListData() []ListItem {
	if m.Tabs.ActiveTab == ideasTab {
		return loadIdeaData()
	}
	if m.Tabs.ActiveTab == projectsTab {
		return loadProjectData()
	}
	if m.Tabs.ActiveTab == booksTab {
		return loadBookData()
	}

	return nil
}

func InitialModel() model {
	ensureDirsExist()
	tabs := []string{
		"Ideas",
		"Projects",
		"Books",
	}

	ideas := loadIdeaData()
	books := loadBookData()
	projects := loadProjectData()

	const defaultWidth = 20

	// Create new list
	items := []list.Item{}
	for _, i := range ideas {
		items = append(items, i)
	}
	ideaDelegate := list.NewDefaultDelegate()
	ideasList := NewList(items, ideaDelegate)

	items = []list.Item{}
	for _, i := range books {
		items = append(items, i)
	}
	bookDelegate := list.NewDefaultDelegate()
	booksList := NewList(items, bookDelegate)

	items = []list.Item{}
	for _, i := range projects {
		items = append(items, i)
	}
	projectDelegate := list.NewDefaultDelegate()
	projectsList := NewList(items, projectDelegate)

	ideavp := viewport.New(0, 0)
	if len(ideas) != 0 {
		content, err := glamour.Render(ideas[0].Description(), "dark")
		if err == nil {
			ideavp.SetContent(content)
		}
	}

	projvp := viewport.New(0, 0)
	if len(projects) != 0 {
		content, err := glamour.Render(projects[0].Description(), "dark")
		if err == nil {
			projvp.SetContent(content)
		}
	}

	bookvp := viewport.New(0, 0)
	if len(books) != 0 {
		content, err := glamour.Render(books[0].Description(), "dark")
		if err == nil {
			bookvp.SetContent(content)
		}
	}
	m := model{
		keys:        keymap.Keys,
		IsTouched:   false,
		activePanel: 1,
		n_panels:    3,
		help:        help.New(),
		theme:       DefaultTheme(),
		Tabs:        components.NewTabModel(tabs, []string{"", ""}, 0),
		IdeaManager: Manager{
			tabIndex:     ideasTab,
			DirPath:      ideasFolder,
			List:         ideasList,
			ListDelegate: ideaDelegate,
			Form:         NewIdeasForm(),
			Viewport:     ideavp,
		},
		ProjectManager: Manager{
			tabIndex:     projectsTab,
			DirPath:      projectsFolder,
			List:         projectsList,
			ListDelegate: projectDelegate,
			Form:         NewForm("New Project"),
			Viewport:     projvp,
		},
		BookManager: Manager{
			tabIndex:     booksTab,
			DirPath:      booksFolder,
			List:         booksList,
			ListDelegate: bookDelegate,
			Form:         NewBooksForm(),
			Viewport:     bookvp,
		},
	}

	return m
}

func (m model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 0)
	cmd := m.Tabs.Init()
	cmds = append(cmds, cmd)

	icmds := m.IdeaManager.Init()
	cmds = append(cmds, icmds...)

	icmds = m.BookManager.Init()
	cmds = append(cmds, icmds...)

	icmds = m.ProjectManager.Init()
	cmds = append(cmds, icmds...)

	cmd = m.BookManager.Form.Init()
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}
